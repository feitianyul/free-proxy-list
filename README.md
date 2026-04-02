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

最后更新：2026-04-02 15:50:46 UTC（2026-04-02 23:50:46 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1704ms | 否 | ✓ 1373ms | ✓ 1351ms | ✓ 934ms | http |
| 167.103.115.102:8800 | ✓ 1236ms | 否 | ✓ 1090ms | ✓ 1940ms | ✓ 1162ms | http |
| 95.213.217.168:52004 | ✓ 765ms | ✓ 1664ms | ✓ 1711ms | 否 | ✓ 1810ms | http |
| 113.160.132.26:8080 | ✓ 1653ms | 否 | 否 | ✓ 1687ms | ✓ 1170ms | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1277ms | ✓ 1924ms | ✓ 1443ms | http |
| 1.231.81.166:3128 | ✓ 1726ms | ✓ 1849ms | 否 | ✓ 1951ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1367ms | 否 | ✓ 1366ms | ✓ 1515ms | ✓ 1410ms | http |
| 211.95.152.50:45046 | ✓ 1226ms | ✓ 1514ms | ✓ 1302ms | ✓ 1485ms | ✓ 1194ms | http |
| 45.167.125.21:999 | ✓ 927ms | ✓ 1545ms | ✓ 1384ms | ✓ 1717ms | ✓ 1683ms | http |
| 217.76.245.80:999 | ✓ 776ms | 否 | 否 | ✓ 1885ms | ✓ 1235ms | http |
| 167.103.31.122:8800 | ✓ 1640ms | 否 | ✓ 1341ms | 否 | ✓ 1758ms | http |
| 222.184.48.251:22222 | ✓ 1086ms | ✓ 1593ms | ✓ 1125ms | 否 | ✓ 1161ms | http |
| 35.225.22.61:80 | ✓ 278ms | ✓ 1667ms | 否 | ✓ 956ms | 否 | http |
| 34.96.238.40:8080 | ✓ 975ms | ✓ 1507ms | ✓ 1286ms | 否 | 否 | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 1019ms | ✓ 1255ms | ✓ 943ms | http |
| 208.87.243.199:7878 | ✓ 897ms | 否 | ✓ 1373ms | ✓ 1064ms | ✓ 774ms | http |
| 147.161.239.240:8800 | ✓ 549ms | 否 | ✓ 1121ms | ✓ 1492ms | ✓ 1385ms | http |
| 128.199.121.61:9090 | ✓ 1494ms | 否 | ✓ 1145ms | ✓ 1207ms | ✓ 1009ms | http |
| 38.145.203.19:8447 | ✓ 1335ms | 否 | ✓ 1970ms | ✓ 872ms | ✓ 1593ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1631ms | ✓ 1752ms | ✓ 1368ms | http |
| 167.103.144.127:8800 | ✓ 1828ms | 否 | ✓ 1727ms | ✓ 1669ms | ✓ 1741ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1253ms | ✓ 1804ms | ✓ 1375ms | ✓ 1046ms | http |
| 120.92.212.16:7890 | ✓ 1716ms | 否 | ✓ 1117ms | 否 | ✓ 1863ms | http |
| 45.12.151.226:2829 | ✓ 749ms | 否 | ✓ 621ms | 否 | ✓ 1333ms | http |
| 203.80.138.81:50000 | ✓ 895ms | ✓ 1628ms | ✓ 1274ms | ✓ 1685ms | ✓ 1226ms | http |
| 209.126.84.232:8888 | ✓ 1637ms | ✓ 1826ms | 否 | 否 | ✓ 1304ms | http |
| 177.234.217.88:999 | 否 | 否 | ✓ 1844ms | ✓ 1790ms | ✓ 1616ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1143ms | 否 | ✓ 990ms | ✓ 836ms | http |
| 45.136.198.40:3128 | ✓ 1356ms | ✓ 1557ms | ✓ 1639ms | ✓ 1618ms | ✓ 1750ms | http |
| 42.96.16.158:1311 | ✓ 1873ms | 否 | ✓ 1194ms | ✓ 1410ms | ✓ 1335ms | http |
| 185.53.46.240:55555 | 否 | ✓ 1758ms | ✓ 1654ms | 否 | ✓ 1666ms | http |
| 120.92.212.16:8890 | ✓ 1143ms | ✓ 1605ms | ✓ 1378ms | 否 | 否 | http |
| 38.34.179.88:8446 | ✓ 487ms | ✓ 1935ms | ✓ 1436ms | ✓ 1175ms | ✓ 880ms | http |
| 101.32.244.83:8080 | ✓ 1094ms | 否 | ✓ 1059ms | ✓ 1334ms | ✓ 1383ms | http |
| 121.43.196.210:8222 | ✓ 1068ms | ✓ 1213ms | ✓ 1004ms | ✓ 1231ms | ✓ 1082ms | http |
| 121.43.196.213:8222 | ✓ 1102ms | 否 | ✓ 921ms | ✓ 1260ms | ✓ 1075ms | http |
| 114.55.226.123:10086 | 否 | ✓ 1563ms | ✓ 1137ms | ✓ 1443ms | ✓ 1194ms | http |
| 181.78.44.63:999 | ✓ 908ms | 否 | ✓ 1631ms | ✓ 1625ms | 否 | http |
| 38.145.203.45:8452 | ✓ 313ms | ✓ 809ms | ✓ 286ms | ✓ 1108ms | 否 | http |
| 38.145.203.46:8448 | ✓ 320ms | ✓ 788ms | ✓ 707ms | 否 | ✓ 814ms | http |
| 104.248.151.93:9090 | 否 | 否 | ✓ 977ms | ✓ 1332ms | ✓ 1022ms | http |
| 146.190.80.158:9090 | ✓ 879ms | 否 | 否 | ✓ 1361ms | ✓ 1034ms | http |
| 128.199.113.85:9090 | 否 | 否 | ✓ 1280ms | ✓ 1653ms | ✓ 1002ms | http |
| 38.145.208.227:8451 | ✓ 922ms | ✓ 1081ms | ✓ 825ms | ✓ 1900ms | ✓ 805ms | http |
| 38.34.179.70:8453 | ✓ 652ms | ✓ 1249ms | ✓ 958ms | ✓ 1151ms | ✓ 1853ms | http |
| 45.136.130.192:8451 | ✓ 596ms | 否 | ✓ 492ms | ✓ 1220ms | ✓ 1328ms | http |
| 45.136.130.193:8444 | ✓ 875ms | 否 | ✓ 459ms | ✓ 1280ms | 否 | http |
| 159.223.71.162:8080 | ✓ 872ms | 否 | 否 | ✓ 1331ms | ✓ 1021ms | http |
| 38.145.208.237:8445 | ✓ 968ms | ✓ 948ms | ✓ 909ms | 否 | ✓ 735ms | http |
| 38.145.203.135:8444 | ✓ 894ms | ✓ 820ms | 否 | ✓ 1941ms | ✓ 1576ms | http |
| 38.34.179.193:8452 | ✓ 609ms | ✓ 1008ms | ✓ 1575ms | 否 | ✓ 815ms | http |
| 38.145.208.220:8448 | ✓ 338ms | ✓ 1674ms | ✓ 1592ms | ✓ 1280ms | ✓ 868ms | http |
| 45.136.130.180:8453 | ✓ 1498ms | 否 | ✓ 1404ms | ✓ 1588ms | ✓ 782ms | http |
| 38.145.218.208:8453 | ✓ 1126ms | 否 | ✓ 1317ms | ✓ 1826ms | 否 | http |
| 38.34.179.185:8445 | ✓ 925ms | 否 | ✓ 1740ms | ✓ 1035ms | 否 | http |
| 45.136.131.59:8450 | ✓ 1550ms | ✓ 1539ms | ✓ 881ms | 否 | ✓ 978ms | http |
| 38.34.179.27:8452 | ✓ 878ms | 否 | ✓ 1987ms | 否 | ✓ 918ms | http |
| 107.174.208.190:3128 | ✓ 280ms | ✓ 898ms | ✓ 844ms | ✓ 1278ms | 否 | http |
| 38.145.220.43:8449 | ✓ 1847ms | ✓ 935ms | ✓ 1666ms | 否 | ✓ 1924ms | http |
| 38.145.220.41:8444 | ✓ 1845ms | ✓ 968ms | ✓ 1616ms | 否 | ✓ 1951ms | http |
| 222.184.48.242:22222 | ✓ 1052ms | ✓ 1432ms | ✓ 1019ms | ✓ 1414ms | ✓ 1966ms | http |
| 5.104.87.17:8051 | ✓ 1499ms | 否 | ✓ 1551ms | 否 | ✓ 1359ms | http |
| 159.223.71.162:443 | ✓ 1472ms | 否 | ✓ 1284ms | ✓ 1248ms | ✓ 1041ms | http |
| 65.108.203.37:18080 | ✓ 747ms | 否 | ✓ 1498ms | 否 | ✓ 1267ms | http |
| 38.242.204.27:3128 | ✓ 456ms | 否 | ✓ 474ms | ✓ 1811ms | ✓ 1449ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1488ms | ✓ 1275ms | ✓ 863ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 975ms | ✓ 1237ms | ✓ 1036ms | http |
| 103.39.51.190:8080 | ✓ 1983ms | 否 | 否 | ✓ 1750ms | ✓ 1577ms | http |
| 38.145.220.102:8453 | 否 | ✓ 1869ms | ✓ 1520ms | ✓ 1740ms | ✓ 1169ms | http |
| 167.160.191.204:6005 | 否 | ✓ 1295ms | ✓ 1482ms | ✓ 1072ms | 否 | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1072ms | ✓ 1227ms | ✓ 1022ms | http |
| 38.34.179.155:8448 | ✓ 456ms | ✓ 1059ms | ✓ 1907ms | ✓ 1189ms | 否 | http |
| 45.140.147.155:1081 | ✓ 1690ms | ✓ 1645ms | ✓ 1501ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 549ms | 否 | ✓ 920ms | ✓ 1645ms | ✓ 1125ms | http |
| 38.145.208.222:8443 | 否 | ✓ 1602ms | ✓ 431ms | ✓ 872ms | ✓ 823ms | http |

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
