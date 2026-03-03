**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
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
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

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

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-03 16:42:53 UTC（2026-03-04 00:42:53 UTC+8）

**代理总数：40**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 40 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 162.240.154.26:3128 | 否 | 否 | ✓ 1519ms | ✓ 1558ms | ✓ 1379ms | http |
| 205.209.118.30:3138 | ✓ 271ms | ✓ 1627ms | ✓ 807ms | ✓ 1053ms | 否 | http |
| 142.171.85.32:1080 | ✓ 981ms | 否 | ✓ 1941ms | ✓ 1807ms | ✓ 1878ms | http |
| 166.0.192.117:8888 | ✓ 329ms | 否 | ✓ 577ms | ✓ 959ms | ✓ 1905ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1422ms | ✓ 1705ms | ✓ 1730ms | http |
| 210.223.44.230:3128 | ✓ 967ms | ✓ 1305ms | ✓ 1643ms | ✓ 1511ms | ✓ 1078ms | http |
| 132.226.235.199:1080 | ✓ 1167ms | 否 | ✓ 1511ms | ✓ 1347ms | ✓ 1064ms | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1745ms | ✓ 1516ms | ✓ 1172ms | http |
| 35.234.17.221:8080 | ✓ 1053ms | 否 | 否 | ✓ 1275ms | ✓ 1192ms | http |
| 81.70.169.194:80 | ✓ 1138ms | ✓ 1475ms | ✓ 1218ms | ✓ 1917ms | 否 | http |
| 162.43.36.42:8080 | ✓ 743ms | 否 | ✓ 965ms | ✓ 1161ms | ✓ 883ms | http |
| 91.193.240.157:9877 | ✓ 781ms | 否 | ✓ 763ms | ✓ 1856ms | ✓ 1970ms | http |
| 201.150.116.32:999 | ✓ 1783ms | 否 | 否 | ✓ 1494ms | ✓ 1284ms | http |
| 45.88.0.113:3128 | ✓ 1446ms | 否 | ✓ 487ms | ✓ 1316ms | 否 | http |
| 45.88.0.99:3128 | ✓ 1458ms | ✓ 1205ms | ✓ 1012ms | ✓ 1589ms | ✓ 1030ms | http |
| 45.88.0.117:3128 | ✓ 470ms | ✓ 1810ms | ✓ 1574ms | ✓ 1278ms | ✓ 1591ms | http |
| 165.101.230.77:8080 | 否 | 否 | ✓ 1542ms | ✓ 1717ms | ✓ 1602ms | http |
| 45.88.0.116:3128 | ✓ 784ms | 否 | 否 | ✓ 1740ms | ✓ 1145ms | http |
| 45.136.198.40:3128 | ✓ 1697ms | ✓ 1437ms | ✓ 1500ms | 否 | ✓ 1491ms | http |
| 120.92.212.16:8890 | ✓ 1135ms | 否 | 否 | ✓ 1465ms | ✓ 1153ms | http |
| 138.124.53.25:7443 | ✓ 1043ms | 否 | ✓ 1713ms | ✓ 1786ms | ✓ 1512ms | http |
| 103.84.95.54:7890 | ✓ 883ms | 否 | 否 | ✓ 1085ms | ✓ 1650ms | http |
| 45.88.0.98:3128 | ✓ 1602ms | 否 | ✓ 797ms | ✓ 1788ms | 否 | http |
| 45.88.0.111:3128 | ✓ 1595ms | ✓ 1985ms | ✓ 812ms | ✓ 1786ms | 否 | http |
| 45.88.0.114:3128 | ✓ 438ms | ✓ 1248ms | ✓ 377ms | ✓ 1267ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1753ms | ✓ 1545ms | ✓ 1461ms | ✓ 1485ms | ✓ 1269ms | http |
| 106.14.203.63:3333 | ✓ 1045ms | ✓ 1280ms | ✓ 1126ms | ✓ 1306ms | ✓ 1080ms | http |
| 45.88.0.115:3128 | ✓ 445ms | ✓ 1233ms | ✓ 369ms | ✓ 1195ms | ✓ 938ms | http |
| 120.92.212.16:7890 | ✓ 1148ms | 否 | 否 | ✓ 1479ms | ✓ 1974ms | http |
| 14.56.107.244:3128 | ✓ 1941ms | 否 | ✓ 1166ms | ✓ 1363ms | ✓ 1508ms | http |
| 144.31.25.69:21064 | ✓ 781ms | 否 | ✓ 1962ms | 否 | ✓ 1484ms | http |
| 121.128.121.54:3128 | ✓ 1104ms | 否 | ✓ 1264ms | ✓ 1176ms | ✓ 1010ms | http |
| 115.231.181.40:8128 | ✓ 1520ms | 否 | 否 | ✓ 1858ms | ✓ 1132ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1074ms | ✓ 1054ms | ✓ 903ms | http |
| 62.113.119.14:8080 | ✓ 710ms | 否 | ✓ 1004ms | ✓ 1719ms | ✓ 1188ms | http |
| 150.249.255.91:3128 | ✓ 1921ms | ✓ 1522ms | 否 | ✓ 1052ms | ✓ 1987ms | http |
| 2.56.178.131:443 | ✓ 760ms | 否 | ✓ 1115ms | ✓ 1905ms | 否 | http |
| 77.83.203.6:443 | ✓ 1936ms | 否 | 否 | ✓ 1767ms | ✓ 1764ms | http |
| 192.71.213.85:9812 | ✓ 393ms | 否 | ✓ 627ms | ✓ 1468ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1913ms | 否 | 否 | ✓ 1693ms | ✓ 1692ms | http |

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
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
