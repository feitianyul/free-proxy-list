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

最后更新：2026-04-15 10:17:36 UTC（2026-04-15 18:17:36 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 628ms | 否 | ✓ 773ms | ✓ 1037ms | ✓ 910ms | http |
| 147.161.239.240:8800 | ✓ 880ms | ✓ 1926ms | ✓ 1127ms | ✓ 1830ms | ✓ 1527ms | http |
| 167.103.115.102:8800 | ✓ 903ms | 否 | ✓ 1202ms | 否 | ✓ 1020ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 922ms | ✓ 1283ms | ✓ 1009ms | http |
| 144.31.27.49:1080 | ✓ 747ms | 否 | ✓ 1652ms | 否 | ✓ 1684ms | http |
| 167.103.34.108:8800 | ✓ 1465ms | 否 | ✓ 1527ms | 否 | ✓ 1502ms | http |
| 167.103.144.127:8800 | ✓ 1669ms | 否 | ✓ 1117ms | ✓ 1693ms | ✓ 1545ms | http |
| 1.231.81.166:3128 | ✓ 1050ms | ✓ 1165ms | ✓ 1253ms | ✓ 1739ms | ✓ 1141ms | http |
| 78.11.96.22:8888 | ✓ 1245ms | ✓ 1789ms | ✓ 1253ms | ✓ 1741ms | ✓ 1872ms | http |
| 167.103.31.122:8800 | ✓ 1346ms | 否 | ✓ 1410ms | ✓ 1763ms | ✓ 1528ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1263ms | ✓ 1083ms | ✓ 1631ms | http |
| 157.120.32.86:3128 | 否 | 否 | ✓ 1649ms | ✓ 1786ms | ✓ 1610ms | http |
| 120.92.108.86:7890 | ✓ 1584ms | 否 | 否 | ✓ 1736ms | ✓ 1441ms | http |
| 59.46.216.131:30001 | ✓ 977ms | 否 | 否 | ✓ 1325ms | ✓ 1140ms | http |
| 12.89.176.82:3128 | ✓ 1113ms | 否 | ✓ 1027ms | ✓ 1255ms | ✓ 958ms | http |
| 157.230.178.216:8088 | ✓ 1125ms | ✓ 1936ms | ✓ 958ms | ✓ 1307ms | ✓ 1291ms | http |
| 210.223.44.230:3128 | ✓ 1303ms | ✓ 1255ms | 否 | ✓ 1563ms | ✓ 1320ms | http |
| 168.110.52.228:3128 | ✓ 552ms | 否 | 否 | ✓ 1884ms | ✓ 677ms | http |
| 152.32.132.190:7890 | ✓ 1798ms | 否 | ✓ 1617ms | ✓ 975ms | 否 | http |
| 185.132.178.178:1080 | ✓ 1031ms | ✓ 1775ms | ✓ 1317ms | 否 | 否 | http |
| 195.26.224.49:3128 | ✓ 1036ms | 否 | ✓ 1086ms | ✓ 1837ms | ✓ 1779ms | http |
| 45.167.125.21:999 | ✓ 1164ms | ✓ 1576ms | ✓ 1542ms | 否 | ✓ 1846ms | http |
| 177.234.217.88:999 | ✓ 1597ms | ✓ 1766ms | ✓ 1958ms | ✓ 1909ms | ✓ 1636ms | http |
| 36.103.198.235:7890 | ✓ 1351ms | ✓ 1373ms | ✓ 1306ms | 否 | 否 | http |
| 147.45.214.210:1080 | ✓ 1027ms | 否 | ✓ 1817ms | 否 | ✓ 1452ms | http |
| 35.225.22.61:80 | ✓ 532ms | 否 | ✓ 833ms | 否 | ✓ 1249ms | http |
| 20.127.128.70:8080 | ✓ 663ms | ✓ 1634ms | ✓ 777ms | ✓ 1355ms | ✓ 1278ms | http |
| 138.124.99.216:8888 | ✓ 1626ms | 否 | ✓ 1767ms | ✓ 1914ms | ✓ 1925ms | http |
| 181.78.44.63:999 | ✓ 1225ms | ✓ 1660ms | ✓ 1423ms | ✓ 1490ms | 否 | http |
| 107.173.42.121:7890 | 否 | ✓ 1125ms | ✓ 769ms | ✓ 1686ms | 否 | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1889ms | ✓ 1426ms | ✓ 1125ms | http |
| 116.80.47.79:3128 | ✓ 1501ms | 否 | ✓ 1472ms | ✓ 1782ms | 否 | http |
| 45.140.147.82:1081 | ✓ 609ms | ✓ 1400ms | ✓ 1324ms | ✓ 1967ms | ✓ 1412ms | http |
| 159.223.225.118:8888 | ✓ 782ms | 否 | 否 | ✓ 1778ms | ✓ 1516ms | http |
| 120.92.211.211:7890 | ✓ 1686ms | ✓ 1816ms | ✓ 1862ms | ✓ 1488ms | 否 | http |
| 83.219.250.8:62920 | ✓ 751ms | ✓ 1991ms | ✓ 1906ms | 否 | 否 | http |
| 88.210.20.237:3128 | ✓ 608ms | ✓ 1612ms | 否 | 否 | ✓ 1546ms | http |
| 109.234.39.202:3128 | 否 | 否 | ✓ 1573ms | ✓ 1511ms | ✓ 1327ms | http |
| 121.230.9.203:1080 | 否 | ✓ 1501ms | ✓ 1187ms | 否 | ✓ 1577ms | http |
| 120.92.212.16:7890 | ✓ 1079ms | 否 | ✓ 1149ms | ✓ 1514ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1360ms | 否 | ✓ 1725ms | ✓ 1612ms | ✓ 1316ms | http |
| 57.128.188.167:9174 | ✓ 1862ms | 否 | ✓ 1785ms | 否 | ✓ 1764ms | http |
| 91.193.240.157:9877 | ✓ 1399ms | 否 | ✓ 1146ms | 否 | ✓ 1855ms | http |
| 36.141.21.200:7890 | ✓ 952ms | ✓ 1278ms | ✓ 983ms | 否 | 否 | http |
| 103.170.196.74:8080 | ✓ 1245ms | 否 | ✓ 1673ms | ✓ 1651ms | ✓ 1294ms | http |
| 140.238.242.189:8100 | ✓ 1106ms | 否 | ✓ 1325ms | ✓ 1992ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1176ms | ✓ 1264ms | ✓ 851ms | 否 | 否 | http |
| 45.149.92.147:5001 | ✓ 935ms | 否 | 否 | ✓ 1243ms | ✓ 716ms | http |
| 103.229.124.12:7890 | ✓ 1401ms | 否 | ✓ 1477ms | ✓ 1576ms | ✓ 1836ms | http |
| 121.138.61.193:8767 | 否 | ✓ 1139ms | ✓ 860ms | ✓ 996ms | ✓ 781ms | http |
| 120.92.212.16:8890 | ✓ 951ms | ✓ 1194ms | ✓ 951ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 363ms | 否 | ✓ 275ms | ✓ 1310ms | ✓ 897ms | http |
| 85.239.59.252:7890 | ✓ 822ms | ✓ 1678ms | ✓ 1466ms | 否 | 否 | http |
| 171.227.167.109:1004 | ✓ 1496ms | 否 | 否 | ✓ 1221ms | ✓ 968ms | http |
| 45.140.147.82:1082 | ✓ 592ms | ✓ 1448ms | ✓ 1544ms | 否 | 否 | http |
| 185.212.119.154:3128 | ✓ 1291ms | 否 | ✓ 1088ms | 否 | ✓ 1903ms | http |
| 202.40.186.66:43773 | ✓ 1654ms | 否 | 否 | ✓ 1988ms | ✓ 1760ms | http |
| 150.107.140.238:3128 | ✓ 1633ms | 否 | ✓ 1910ms | ✓ 1236ms | 否 | http |
| 171.227.167.109:1006 | ✓ 1504ms | 否 | ✓ 1077ms | ✓ 1236ms | ✓ 1265ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1987ms | ✓ 1880ms | ✓ 1500ms | http |
| 77.110.113.24:40000 | ✓ 1369ms | 否 | ✓ 1541ms | 否 | ✓ 1998ms | http |
| 121.230.8.247:1080 | ✓ 1414ms | ✓ 1562ms | 否 | ✓ 1597ms | ✓ 1406ms | http |
| 121.230.8.153:1080 | ✓ 941ms | ✓ 1443ms | ✓ 1605ms | ✓ 1305ms | ✓ 1026ms | http |
| 101.32.243.189:80 | ✓ 1352ms | 否 | ✓ 1821ms | 否 | ✓ 1242ms | http |
| 121.230.8.253:1080 | ✓ 1191ms | ✓ 1670ms | ✓ 1590ms | 否 | 否 | http |
| 121.230.9.198:1080 | ✓ 1869ms | 否 | ✓ 1070ms | ✓ 1486ms | 否 | http |
| 61.52.131.172:8443 | ✓ 907ms | ✓ 1071ms | ✓ 953ms | ✓ 1248ms | ✓ 925ms | http |
| 144.31.25.69:21064 | ✓ 1199ms | 否 | ✓ 1212ms | 否 | ✓ 1833ms | http |
| 103.122.65.242:8080 | 否 | 否 | ✓ 1797ms | ✓ 1833ms | ✓ 1561ms | http |

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
