**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/gfpcom/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/gfpcom/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/gfpcom/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：对每条 HTTP/HTTPS 代理访问以下两个地址进行验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://finance.sina.com.cn/`（新浪财经）
  两个请求均需在 **2 秒内**成功（HTTP 200）方视为通过，未通过的不写入列表。校验时**多代理并发**、**单代理内双 URL 并行**，以提升吞吐。每个代理会分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次，并生成「代理地址 | HTTP | HTTPS」结果表格与 `http+s.txt`（同时支持两种协议的代理）。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=2000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 2000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。

<!-- BEGIN PROXY LIST -->

最后更新：2026-02-16 22:51:14 UTC（2026-02-17 06:51:14 UTC+8）

**代理总数：102**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 102 | https://raw.githubusercontent.com/wiki/gfpcom/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/gfpcom/free-proxy-list/lists/https.txt |
| HTTP+S | 0 | https://raw.githubusercontent.com/wiki/gfpcom/free-proxy-list/lists/http+s.txt |

<!-- END PROXY LIST -->

以下为最近校验结果的前 100 条（代理地址 | HTTP | HTTPS），便于查看双协议可用性。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | HTTP | HTTPS |
|----------|------|--------|
| 138.68.60.8:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 67.43.228.250:26419 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 67.43.228.250:24763 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 72.10.160.90:27113 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 67.43.228.254:28073 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 134.209.29.120:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 47.74.226.8:5001 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 128.199.202.122:8080 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 72.56.59.62:63133 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 72.56.50.17:59787 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.30.86:63900 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.30.68:63744 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.30.58:63744 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.30.63:63744 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.30.91:63900 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.30.37:59741 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 72.56.59.56:63127 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 205.209.118.30:3138 | ✓ 1196ms | 否 Head "https://www.ea |
| 5.9.218.168:3128 | 否 Head "https://financ | 否 Head "https://www.ea |
| 158.160.215.167:8123 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 147.75.202.36:10016 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 94.176.3.43:7443 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 5.75.201.136:1080 | ✓ 1895ms | 否 Head "https://www.ea |
| 45.151.182.9:3128 | 否 Head "https://financ | 否 Head "https://www.ea |
| 213.219.215.233:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 34.101.184.164:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 150.230.104.3:16728 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 140.227.61.201:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 195.158.8.123:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 188.130.160.209:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 77.110.125.30:65531 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 77.110.97.209:8080 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 77.225.198.220:9812 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 90.84.188.97:8000 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 94.176.3.53:7443 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 72.56.59.17:61931 | 否 Head "https://financ | 否 Head "https://www.ea |
| 72.56.59.23:61937 | 否 Head "https://financ | 否 Head "https://www.ea |
| 91.187.57.109:8080 | 否 Head "https://financ | 否 Head "https://www.ea |
| 35.244.232.197:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 45.138.69.40:20000 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 154.31.200.185:6666 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 191.101.1.116:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 45.138.69.35:20000 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 192.73.244.36:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 149.28.195.13:32769 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.229:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.16.0.104:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 43.153.73.106:443 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.85:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 132.148.6.255:20859 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.160:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 43.153.64.66:443 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.219:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 166.62.35.102:45775 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.225:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.238.100.115:45314 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.199:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.116:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.223.45.93:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.169:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 185.18.250.181:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.50:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 185.18.250.232:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 203.32.120.91:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 185.18.250.83:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.231:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.119:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.66:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.105:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.64:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 132.148.244.30:45157 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.122:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.10:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.161:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.185:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.69:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.42:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.101:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 192.154.246.96:9000 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 192.163.200.196:24787 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 173.245.49.173:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 128.199.1.8:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.129.203.245:10733 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 162.241.129.84:36504 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.129.202.127:12354 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 45.189.252.93:999 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 209.145.60.213:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 5.78.86.132:8080 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 45.56.112.189:7497 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 162.240.212.185:59991 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 204.252.113.141:3128 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.129.203.245:10026 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 64.227.106.157:80 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.129.202.127:10810 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 104.129.203.245:10139 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 66.63.168.119:8000 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 107.174.231.218:8888 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 50.116.29.136:32728 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 71.19.249.97:8443 | 否 Head "https://www.ea | 否 Head "https://www.ea |
| 162.144.83.73:60468 | 否 Head "https://www.ea | 否 Head "https://www.ea |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=gfpcom/free-proxy-list&type=Date)](https://star-history.com/#gfpcom/free-proxy-list&Date)
