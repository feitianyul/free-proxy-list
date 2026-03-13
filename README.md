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

最后更新：2026-03-13 19:43:31 UTC（2026-03-14 03:43:31 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 854ms | ✓ 978ms | ✓ 773ms | ✓ 831ms | ✓ 693ms | http |
| 205.209.118.30:3138 | ✓ 874ms | ✓ 1892ms | ✓ 738ms | ✓ 1211ms | ✓ 889ms | http |
| 43.167.227.161:1080 | ✓ 1919ms | ✓ 1968ms | 否 | ✓ 930ms | ✓ 1199ms | http |
| 113.160.132.26:8080 | ✓ 1674ms | ✓ 1392ms | 否 | 否 | ✓ 1235ms | http |
| 103.87.67.75:3129 | 否 | 否 | ✓ 1661ms | ✓ 1708ms | ✓ 1939ms | http |
| 45.167.124.52:8080 | ✓ 1569ms | 否 | ✓ 1307ms | ✓ 1966ms | ✓ 1553ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1554ms | ✓ 1129ms | 否 | ✓ 1043ms | http |
| 8.219.97.248:80 | ✓ 1492ms | 否 | ✓ 1146ms | ✓ 1908ms | 否 | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 802ms | ✓ 1135ms | ✓ 977ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1339ms | ✓ 1573ms | ✓ 1542ms | http |
| 186.148.180.46:999 | ✓ 1262ms | 否 | ✓ 1320ms | ✓ 1588ms | ✓ 1335ms | http |
| 45.140.147.82:1081 | ✓ 836ms | ✓ 1931ms | ✓ 1364ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 1193ms | ✓ 1248ms | 否 | ✓ 1138ms | 否 | http |
| 101.43.255.96:80 | ✓ 1131ms | ✓ 1450ms | ✓ 1089ms | ✓ 1389ms | ✓ 1215ms | http |
| 81.70.169.194:80 | ✓ 1125ms | 否 | ✓ 1065ms | ✓ 1295ms | ✓ 1110ms | http |
| 45.136.130.175:8443 | ✓ 1800ms | 否 | ✓ 1990ms | ✓ 999ms | 否 | http |
| 171.251.173.39:5104 | ✓ 1502ms | 否 | ✓ 1731ms | ✓ 1654ms | ✓ 1606ms | http |
| 123.57.0.163:8888 | 否 | ✓ 1519ms | 否 | ✓ 1507ms | ✓ 1509ms | http |
| 45.136.131.47:8443 | ✓ 684ms | ✓ 909ms | ✓ 189ms | ✓ 1204ms | ✓ 612ms | http |
| 178.236.245.59:3128 | ✓ 679ms | ✓ 1697ms | ✓ 833ms | ✓ 1647ms | ✓ 1322ms | http |
| 103.84.95.54:7890 | ✓ 1139ms | 否 | 否 | ✓ 1952ms | ✓ 1126ms | http |
| 120.92.212.16:7890 | ✓ 1107ms | ✓ 1621ms | 否 | ✓ 1557ms | 否 | http |
| 178.236.245.17:3128 | ✓ 1939ms | 否 | ✓ 772ms | ✓ 1833ms | ✓ 1845ms | http |
| 185.191.236.162:3128 | ✓ 1779ms | 否 | ✓ 732ms | ✓ 1639ms | ✓ 1159ms | http |
| 91.247.126.241:2080 | 否 | 否 | ✓ 928ms | ✓ 1992ms | ✓ 1869ms | http |
| 121.126.185.63:25152 | 否 | ✓ 1944ms | ✓ 1388ms | 否 | ✓ 1944ms | http |
| 86.53.183.16:1080 | ✓ 1044ms | 否 | ✓ 1338ms | 否 | ✓ 1537ms | http |
| 162.240.154.26:3128 | ✓ 637ms | 否 | ✓ 1587ms | ✓ 1551ms | ✓ 845ms | http |
| 165.227.5.10:8888 | ✓ 644ms | ✓ 1553ms | ✓ 845ms | 否 | 否 | http |
| 64.188.90.36:1080 | ✓ 1694ms | 否 | ✓ 854ms | ✓ 1715ms | 否 | http |
| 121.230.8.111:1080 | 否 | ✓ 1879ms | ✓ 1090ms | ✓ 1669ms | ✓ 1176ms | http |
| 150.230.249.50:1080 | ✓ 1719ms | 否 | ✓ 1282ms | 否 | ✓ 1872ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1402ms | 否 | ✓ 1493ms | ✓ 1188ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1616ms | ✓ 1151ms | ✓ 959ms | http |
| 116.80.49.162:3172 | 否 | 否 | ✓ 1758ms | ✓ 1906ms | ✓ 1737ms | http |
| 45.136.198.40:3128 | ✓ 717ms | 否 | ✓ 807ms | ✓ 1696ms | ✓ 1423ms | http |
| 47.101.149.27:9010 | ✓ 1393ms | ✓ 1596ms | ✓ 1847ms | ✓ 1319ms | ✓ 1425ms | http |
| 210.223.44.230:3128 | ✓ 1791ms | 否 | 否 | ✓ 1743ms | ✓ 1801ms | http |
| 201.144.20.238:3128 | ✓ 1899ms | ✓ 1760ms | ✓ 1220ms | ✓ 1442ms | ✓ 1042ms | http |
| 106.117.208.101:7890 | ✓ 1041ms | ✓ 1420ms | ✓ 1095ms | ✓ 1414ms | ✓ 1043ms | http |
| 42.84.157.30:10808 | ✓ 1775ms | ✓ 1368ms | ✓ 1180ms | ✓ 1412ms | ✓ 1290ms | http |
| 152.70.84.108:8080 | ✓ 1538ms | ✓ 1968ms | ✓ 1440ms | ✓ 992ms | ✓ 1425ms | http |
| 152.42.213.210:443 | ✓ 900ms | 否 | ✓ 1641ms | ✓ 1174ms | ✓ 936ms | http |
| 45.136.130.223:8443 | ✓ 373ms | ✓ 945ms | ✓ 407ms | ✓ 850ms | ✓ 646ms | http |
| 61.76.95.217:40088 | ✓ 1996ms | ✓ 1783ms | ✓ 1493ms | ✓ 1374ms | ✓ 1074ms | http |
| 103.39.51.190:8080 | ✓ 1844ms | 否 | 否 | ✓ 1514ms | ✓ 1418ms | http |
| 62.113.119.14:8080 | ✓ 1366ms | ✓ 1623ms | ✓ 1466ms | ✓ 1778ms | ✓ 1190ms | http |
| 24.199.124.152:3128 | 否 | ✓ 1407ms | ✓ 1025ms | ✓ 837ms | ✓ 622ms | http |
| 128.199.114.189:9090 | ✓ 1053ms | 否 | ✓ 1556ms | ✓ 1463ms | ✓ 1728ms | http |
| 128.199.120.45:9090 | 否 | 否 | ✓ 1829ms | ✓ 1671ms | ✓ 1010ms | http |
| 162.248.165.72:1080 | ✓ 1153ms | ✓ 1748ms | ✓ 917ms | 否 | 否 | http |
| 128.199.254.13:9090 | ✓ 1680ms | 否 | ✓ 1647ms | ✓ 1481ms | ✓ 1002ms | http |
| 61.52.131.172:8443 | ✓ 1161ms | ✓ 1288ms | ✓ 996ms | ✓ 1262ms | ✓ 1007ms | http |
| 172.212.68.37:3128 | ✓ 426ms | ✓ 1541ms | ✓ 1765ms | ✓ 1545ms | ✓ 1110ms | http |
| 104.129.203.245:10026 | ✓ 820ms | ✓ 1046ms | ✓ 720ms | ✓ 1074ms | ✓ 670ms | http |
| 104.129.203.245:10733 | ✓ 821ms | ✓ 1023ms | ✓ 701ms | ✓ 1020ms | ✓ 775ms | http |
| 104.129.203.245:10139 | ✓ 820ms | ✓ 878ms | ✓ 889ms | ✓ 1038ms | ✓ 1008ms | http |
| 45.168.238.193:8443 | ✓ 817ms | ✓ 1074ms | ✓ 1118ms | ✓ 1014ms | ✓ 885ms | http |
| 47.105.98.23:3128 | ✓ 1171ms | ✓ 1265ms | ✓ 1463ms | ✓ 1321ms | ✓ 1231ms | http |

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
