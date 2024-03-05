1. docker network inspect my-large-scale-system_default
   1. 查看 container 的網路資料
2. 避免每次都將 images 刪除，重新 docker compose up -d
   1. docker-compose up --build
      1. 全服務重新 build
   2. docker compose build <服務名稱>


docker network
* 为什么要使用 Docker 网络？
  * 目前所有组件都在云主机上运行，并使用主机网络。
  * 组件之间通过主机名进行通信 (例如，web 应用调用授权服务时使用 localhost)。
  * 这并不是真实系统的配置，实际系统使用的是真实主机或虚拟机，拥有各自的主机名。
* 解决方案：
  * 创建一个名为 "docker-network" 的 Docker 桥接网络。
  * 将所有容器加入到该网络中。
  * 组件之间可以使用各自的容器名进行通信，而无需再使用 localhost。
* 为什么要使用 Docker 卷？
  * 容器中的数据会保存在容器内部的文件系统中
  * 停止或删除容器后，这些数据也会丢失，因为容器是短暂的。
* 解决方案：
  * 在主机上创建 Docker 卷。
  * 将容器中需要持久保存数据的目录挂载到 Docker 卷。
  * 这样，数据就保存在主机上，由 Docker 进行管理，即使容器被删除，数据也不会丢失。


* 當更改 docker compose.yml，需要刪除 container ，也需要刪除 images ，才可以套用到新的設定
  * docker-compose down
  * docker-compose rm -f
  * 如果是更改網路設定，可以用 restart
    * docker-compose restart <container_name>

* gateway
  *  本文讨论了微服务架构中路由问题的解决方案，重点介绍了使用 Netflix Zuul 网关服务进行服务路由管理。
  *  1. 问题：微服务架构中，客户端（例如单页应用、移动应用）需要知道服务的位置（主机、端口）才能发起请求。
  *  这种方式使得客户端配置复杂，且服务配置变更时需要更新所有客户端。
  *  2. 解决方案：网关服务
  *  引入网关服务作为中间层，所有请求都先经过网关服务。
  *  网关服务根据 URL 等信息将请求路由到相应的微服务。
  *  客户端只需要知道网关服务的地址，无需关心后端微服务的具体位置。
  *  3. 实现方式：
  *  可以使用反向代理 (例如 Nginx) 实现网关服务的功能，但功能有限。
  *  推荐使用 Netflix Zuul 网关服务：
  *  与 Spring Cloud 框架兼容，便于集成。
  *  提供丰富的功能，例如:
  *  集中式认证
  *  请求/响应日志记录
* "logs"（日誌）、"trace"（追蹤）、和 "metrics"（指標或性能指標） 是三個常見的術語，它們用於監控和管理應用程式的不同方面。以下是對這些術語的簡單解釋：
* 		Logs（日誌）：
    * 日誌是應用程式或系統在運行時生成的詳細記錄。這些記錄通常包含有關應用程式的事件、錯誤、警告和其他重要信息的記錄。
    * 日誌對於故障排除、調試和了解應用程式運行狀態非常重要。它們可以提供對應用程式內部運作的洞察，有助於識別問題和追蹤應用程式的行為。
    * 日志收集:
      * 在每個生成日誌的機器上部署日志代理 (agent)。代理實時讀取日誌並將其發送到日志收集器 (log collector)。日志处理和存储:日誌收集器接收來自各個代理的日誌。收集器處理日誌，使其具有可分析的結構。將處理后的日誌永久存儲到可靠的存儲系統中。日志分析:分析组件讀取存儲的日誌信息。分析组件可以生成報告和其他形式的分析結果。
    * 日志收集:
      * 選擇 Fluentd 作為日志代理，將各個機器上的日誌傳輸到日志收集器。對於 Docker 容器化的組件，可以使用 Fluentd 的 Docker 驅動程序，無需額外安裝代理。日志存储:選擇 Elasticsearch 作為日志存儲，因為它適合高吞吐量數據存儲和检索，且與 Fluentd 集成良好。日志分析:選擇 Kibana 作為日志分析工具，因为它易於使用，可以提供可视化的界面和报告。
* 		Trace（追蹤）：
    * 追蹤是用於追蹤應用程式中不同組件之間的操作流程的工具。它允許開發者了解一個請求或事務如何在系統中傳播，並且可以標示各個組件的性能。
    * 追蹤可以用來識別性能瓶頸、優化代碼以及理解複雜的系統互動。分佈式系統中的追蹤對於解決跨越多個服務的性能問題尤為重要。
    * 查看延遲時間
* 		Metrics（指標）：
    * 指標是用於量化和衡量應用程式或系統性能的數據。這些數據可以包括各種方面，例如請求速率、錯誤率、CPU使用率等。
    * 指標通常用於監控應用程式的運行狀態，以及評估性能和健康狀態。它們提供了對系統行為的定量洞察，有助於預測和防止問題。
    * 吞吐量 throughput


