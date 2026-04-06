package main

import (
	"fmt"
	"log"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("Lua 字节码功能测试")
	fmt.Println("========================================")

	// 初始化 Lua 引擎
	engine := lua_engine.NewLuaEngine(nil)
	defer engine.Close()

	// 测试 1: 编译 Lua 源码字符串为字节码
	fmt.Println("\n测试 1: 编译 Lua 源码字符串为字节码")
	fmt.Println("----------------------------------------")
	
	source := `
		-- 这是一个简单的 Lua 脚本
		local function add(a, b)
			return a + b
		end
		
		local function multiply(a, b)
			return a * b
		end
		
		-- 测试函数调用
		print("测试加法: 3 + 5 = " .. add(3, 5))
		print("测试乘法: 4 * 6 = " .. multiply(4, 6))
		
		-- 返回结果
		return "字节码测试成功!"
	`
	
	// 编译源码为字节码
	bytecode, err := engine.CompileString(source, "test_script")
	if err != nil {
		log.Fatalf("编译字节码失败: %v", err)
	}
	
	fmt.Printf("字节码编译成功! 名称: %s\n", bytecode.GetName())
	fmt.Printf("函数原型: %v\n", bytecode.GetFunctionProto() != nil)

	// 测试 2: 执行编译后的字节码
	fmt.Println("\n测试 2: 执行编译后的字节码")
	fmt.Println("----------------------------------------")
	
	err = engine.ExecuteBytecode(bytecode)
	if err != nil {
		log.Fatalf("执行字节码失败: %v", err)
	}
	
	fmt.Println("字节码执行成功!")

	// 测试 3: 多次执行同一字节码（验证缓存效果）
	fmt.Println("\n测试 3: 多次执行同一字节码")
	fmt.Println("----------------------------------------")
	
	for i := 0; i < 3; i++ {
		fmt.Printf("第 %d 次执行:\n", i+1)
		err = engine.ExecuteBytecode(bytecode)
		if err != nil {
			log.Fatalf("第 %d 次执行字节码失败: %v", i+1, err)
		}
	}

	// 测试 4: 编译复杂脚本
	fmt.Println("\n测试 4: 编译复杂脚本")
	fmt.Println("----------------------------------------")
	
	complexSource := `
		-- 复杂脚本测试
		local M = {}
		
		function M.factorial(n)
			if n <= 1 then
				return 1
			end
			return n * M.factorial(n - 1)
		end
		
		function M.fibonacci(n)
			if n <= 2 then
				return 1
			end
			local a, b = 1, 1
			for i = 3, n do
				a, b = b, a + b
			end
			return b
		end
		
		-- 测试函数
		print("阶乘测试:")
		for i = 1, 5 do
			print(string.format("  %d! = %d", i, M.factorial(i)))
		end
		
		print("\n斐波那契测试:")
		for i = 1, 10 do
			print(string.format("  fib(%d) = %d", i, M.fibonacci(i)))
		end
		
		return M
	`
	
	complexBytecode, err := engine.CompileString(complexSource, "complex_script")
	if err != nil {
		log.Fatalf("编译复杂脚本失败: %v", err)
	}
	
	fmt.Println("复杂脚本编译成功!")
	
	err = engine.ExecuteBytecode(complexBytecode)
	if err != nil {
		log.Fatalf("执行复杂脚本失败: %v", err)
	}
	
	fmt.Println("复杂脚本执行成功!")

	// 测试 5: 全局函数测试
	fmt.Println("\n测试 5: 全局函数测试")
	fmt.Println("----------------------------------------")
	
	// 使用全局引擎
	globalEngine := lua_engine.GetLuaEngine()
	defer globalEngine.Close()
	
	// 使用全局编译函数
	globalBytecode, err := lua_engine.CompileString("print('全局引擎字节码测试成功!')", "global_test")
	if err != nil {
		log.Fatalf("全局编译失败: %v", err)
	}
	
	// 使用全局执行函数
	err = lua_engine.ExecuteBytecode(globalBytecode)
	if err != nil {
		log.Fatalf("全局执行失败: %v", err)
	}

	fmt.Println("\n========================================")
	fmt.Println("所有测试完成!")
	fmt.Println("========================================")
	
	fmt.Println("\n重要提示:")
	fmt.Println("1. gopher-lua 的字节码格式与标准 Lua 字节码不兼容")
	fmt.Println("2. 字节码仅能在 gopher-lua 虚拟机中执行")
	fmt.Println("3. 预编译字节码可以提高重复执行的效率")
	fmt.Println("4. require 函数现在支持加载 .gluac 字节码文件")
}
