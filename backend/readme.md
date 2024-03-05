1. 啟動 docker 應用程式背景執行
2. docker compose up -d
   1. 將 docker compose 提及的服務啟動起來
3. docker exec -it stock1102 sh
   1. 進入到 docker container 裡面
4. go run main.go
   1. 啟動 該程式進入檔


要注意，當這個資料夾的 go mod init <package-name>
這個 <package-name> 會影響到 import 的時候的名稱