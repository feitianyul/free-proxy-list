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

最后更新：2026-03-06 05:40:31 UTC（2026-03-06 13:40:31 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 430ms | 否 | ✓ 1117ms | ✓ 1348ms | ✓ 1215ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1696ms | ✓ 983ms | ✓ 1278ms | ✓ 743ms | http |
| 103.166.185.54:3128 | ✓ 1380ms | 否 | ✓ 1226ms | ✓ 1594ms | ✓ 919ms | http |
| 91.107.175.112:10801 | ✓ 1126ms | 否 | ✓ 1413ms | 否 | ✓ 1752ms | http |
| 107.174.80.186:3128 | ✓ 724ms | 否 | ✓ 831ms | ✓ 906ms | ✓ 909ms | http |
| 35.225.22.61:80 | ✓ 666ms | 否 | ✓ 1057ms | ✓ 1209ms | ✓ 1140ms | http |
| 61.72.110.54:3128 | ✓ 590ms | ✓ 1260ms | ✓ 1396ms | ✓ 1412ms | ✓ 758ms | http |
| 14.56.107.244:3128 | ✓ 625ms | ✓ 1443ms | ✓ 617ms | ✓ 1032ms | ✓ 1807ms | http |
| 154.37.208.132:30000 | ✓ 1847ms | 否 | 否 | ✓ 1885ms | ✓ 1930ms | http |
| 103.84.95.54:7890 | ✓ 1186ms | 否 | ✓ 1061ms | 否 | ✓ 681ms | http |
| 61.72.110.94:3128 | ✓ 1939ms | 否 | ✓ 805ms | ✓ 1536ms | ✓ 1962ms | http |
| 125.128.12.14:3128 | ✓ 636ms | 否 | 否 | ✓ 1548ms | ✓ 1343ms | http |
| 121.128.121.54:3128 | ✓ 1547ms | 否 | ✓ 743ms | 否 | ✓ 864ms | http |
| 35.234.17.221:8080 | ✓ 1413ms | ✓ 1266ms | 否 | 否 | ✓ 1048ms | http |
| 61.72.221.194:3128 | ✓ 1792ms | ✓ 1231ms | ✓ 922ms | ✓ 1590ms | ✓ 1231ms | http |
| 61.72.221.234:3128 | ✓ 1395ms | ✓ 1278ms | ✓ 1719ms | ✓ 974ms | 否 | http |
| 116.80.82.224:3172 | ✓ 1748ms | 否 | 否 | ✓ 1832ms | ✓ 1641ms | http |
| 116.80.82.216:3172 | ✓ 1748ms | 否 | ✓ 1467ms | ✓ 1801ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1223ms | ✓ 876ms | ✓ 1816ms | 否 | http |
| 183.237.195.130:3128 | ✓ 902ms | ✓ 1270ms | ✓ 1441ms | ✓ 1174ms | ✓ 946ms | http |
| 61.72.221.94:3128 | ✓ 1341ms | 否 | 否 | ✓ 1957ms | ✓ 1856ms | http |
| 81.70.169.194:80 | ✓ 1866ms | ✓ 1265ms | 否 | ✓ 1270ms | 否 | http |
| 91.193.240.157:9877 | ✓ 1430ms | 否 | ✓ 1526ms | 否 | ✓ 1637ms | http |
| 121.126.185.63:25152 | ✓ 1850ms | 否 | ✓ 1751ms | 否 | ✓ 1943ms | http |
| 116.80.82.220:3172 | ✓ 1685ms | 否 | 否 | ✓ 1843ms | ✓ 1621ms | http |
| 14.56.177.44:3128 | ✓ 1421ms | ✓ 1304ms | ✓ 1049ms | ✓ 960ms | ✓ 806ms | http |
| 192.166.82.55:1080 | ✓ 819ms | 否 | ✓ 1420ms | ✓ 1154ms | ✓ 1474ms | http |
| 62.113.119.14:8080 | ✓ 1288ms | 否 | ✓ 862ms | ✓ 1696ms | ✓ 1261ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1521ms | ✓ 1388ms | ✓ 1061ms | http |
| 120.92.212.16:8890 | ✓ 954ms | 否 | 否 | ✓ 1145ms | ✓ 1974ms | http |
| 115.231.181.40:8128 | ✓ 1901ms | 否 | ✓ 870ms | ✓ 1384ms | 否 | http |
| 23.95.76.201:8443 | ✓ 533ms | 否 | ✓ 1955ms | ✓ 1402ms | ✓ 1144ms | http |
| 160.250.4.245:1 | ✓ 1835ms | 否 | ✓ 1594ms | ✓ 1212ms | ✓ 945ms | http |
| 103.215.36.88:19672 | ✓ 1032ms | ✓ 1286ms | ✓ 1385ms | ✓ 1406ms | 否 | http |
| 160.250.5.22:1 | ✓ 1310ms | 否 | ✓ 1549ms | 否 | ✓ 1022ms | http |
| 116.80.82.222:3172 | ✓ 1497ms | 否 | ✓ 1493ms | ✓ 1787ms | 否 | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 521ms | ✓ 884ms | ✓ 582ms | http |
| 138.124.53.25:7443 | ✓ 1491ms | 否 | ✓ 1733ms | ✓ 1727ms | ✓ 1571ms | http |
| 69.48.179.20:3128 | ✓ 806ms | 否 | ✓ 286ms | ✓ 1342ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1549ms | ✓ 1168ms | ✓ 958ms | ✓ 1180ms | ✓ 937ms | http |
| 211.171.114.154:3128 | ✓ 1505ms | 否 | ✓ 1854ms | 否 | ✓ 1541ms | http |
| 103.74.192.243:7890 | 否 | ✓ 1947ms | 否 | ✓ 1996ms | ✓ 1001ms | http |
| 170.78.208.251:999 | ✓ 1837ms | 否 | ✓ 1420ms | ✓ 1478ms | 否 | http |
| 210.223.44.230:3128 | ✓ 843ms | 否 | ✓ 992ms | ✓ 1145ms | ✓ 1968ms | http |
| 45.136.198.40:3128 | ✓ 909ms | 否 | ✓ 1913ms | 否 | ✓ 1990ms | http |
| 46.249.103.192:443 | ✓ 1323ms | 否 | ✓ 812ms | ✓ 1744ms | 否 | http |
| 185.191.236.162:3128 | ✓ 791ms | ✓ 1940ms | ✓ 1939ms | 否 | 否 | http |
| 209.38.54.154:8443 | 否 | 否 | ✓ 1849ms | ✓ 1959ms | ✓ 1643ms | http |
| 5.75.196.26:40000 | ✓ 667ms | 否 | ✓ 1237ms | 否 | ✓ 1942ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1310ms | ✓ 1530ms | ✓ 1771ms | http |
| 47.77.193.180:1080 | ✓ 120ms | ✓ 859ms | ✓ 278ms | ✓ 693ms | ✓ 539ms | http |
| 160.250.4.13:1 | 否 | 否 | ✓ 1597ms | ✓ 1769ms | ✓ 997ms | http |
| 103.215.36.88:17932 | ✓ 1698ms | 否 | ✓ 1246ms | 否 | ✓ 1194ms | http |
| 8.219.97.248:80 | ✓ 1042ms | 否 | ✓ 1226ms | 否 | ✓ 1176ms | http |
| 172.212.68.37:3128 | ✓ 697ms | 否 | ✓ 721ms | ✓ 1681ms | ✓ 1291ms | http |
| 91.233.223.147:3128 | ✓ 1033ms | 否 | ✓ 1346ms | 否 | ✓ 1685ms | http |
| 154.90.48.209:9090 | ✓ 1897ms | 否 | ✓ 1069ms | ✓ 1197ms | ✓ 1019ms | http |
| 74.48.78.224:2080 | ✓ 367ms | ✓ 895ms | ✓ 1504ms | 否 | 否 | http |
| 121.230.9.248:1080 | 否 | ✓ 1430ms | ✓ 1075ms | ✓ 1483ms | ✓ 1467ms | http |
| 116.80.63.64:7777 | ✓ 1510ms | 否 | ✓ 1977ms | ✓ 1981ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1056ms | 否 | ✓ 1161ms | 否 | ✓ 1021ms | http |
| 121.230.8.208:1080 | 否 | 否 | ✓ 1050ms | ✓ 1630ms | ✓ 1151ms | http |
| 57.128.188.167:9174 | ✓ 1616ms | 否 | ✓ 1669ms | ✓ 1948ms | ✓ 1624ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1350ms | 否 | ✓ 1264ms | ✓ 1009ms | http |

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
