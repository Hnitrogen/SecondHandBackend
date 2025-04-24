package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func runInTerminal(title, command string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// PowerShell 方案（推荐）
		psCommand := fmt.Sprintf(`
			$host.UI.RawUI.WindowTitle = "%s"
			%s
			Write-Host "按 Enter 关闭窗口..." -ForegroundColor Yellow
			$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
		`, title, command)

		// 保存为临时 .ps1 文件并执行
		psScript := filepath.Join(os.TempDir(), title+".ps1")
		err := os.WriteFile(psScript, []byte(psCommand), 0644)
		if err != nil {
			fmt.Printf("创建 PowerShell 脚本失败: %v\n", err)
			return
		}

		cmd = exec.Command("powershell", "-NoExit", "-File", psScript)

	case "darwin":
		cmd = exec.Command("osascript", "-e", fmt.Sprintf(
			`tell app "Terminal" to do script "%s"`, command,
		))
	case "linux":
		cmd = exec.Command("gnome-terminal", "--title", title, "--", "bash", "-c", command)
	default:
		fmt.Printf("Unsupported OS: %s\n", runtime.GOOS)
		return
	}

	// 启动进程（不等待结束）
	if err := cmd.Start(); err != nil {
		fmt.Printf("启动 %s 失败: %v\n", title, err)
	}
}

func main() {
	// 启动 Consul（PowerShell 会保持窗口打开）
	runInTerminal("Consul", "consul agent -dev")
	fmt.Println("🚀 Consul已启动！")

	// 启动 Nginx（示例）
	runInTerminal("Nginx", "cd \"C:\\Users\\17914\\Downloads\\tutu\\nginx-1.26.3\"; nginx -c \"C:\\Users\\17914\\Downloads\\tutu\\nginx-1.26.3\\conf\\nginx.conf\"")
	fmt.Println("🚀 Nginx已启动！")

	//// 启动 前端
	//runInTerminal("Web", "cd C:\\Users\\17914\\Desktop\\毕设\\lj\\SecondHandBackend\\web; npm run dev")

	// 启动 Go 微服务（示例）
	// runInTerminal("UserService", "go run ./cmd/user_service/main.go --port 5001")

	fmt.Println("🚀 所有服务已启动！")

	// 等待 Ctrl+C 退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("🛑 收到终止信号，程序退出")
}
