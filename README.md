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

最后更新：2026-03-03 07:51:02 UTC（2026-03-03 15:51:02 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 70ms | 否 | ✓ 975ms | ✓ 889ms | ✓ 675ms | http |
| 160.238.65.5:3128 | 否 | 否 | ✓ 1322ms | ✓ 1577ms | ✓ 1075ms | http |
| 183.249.5.109:22222 | ✓ 923ms | ✓ 1180ms | ✓ 921ms | ✓ 1241ms | ✓ 923ms | http |
| 45.140.147.82:1081 | ✓ 409ms | ✓ 1351ms | ✓ 835ms | 否 | 否 | http |
| 183.249.5.111:22222 | ✓ 939ms | ✓ 1176ms | ✓ 885ms | ✓ 1183ms | ✓ 1379ms | http |
| 2.56.178.131:443 | ✓ 1275ms | 否 | ✓ 1677ms | ✓ 1542ms | 否 | http |
| 74.208.234.198:443 | ✓ 786ms | 否 | ✓ 953ms | ✓ 1770ms | 否 | http |
| 83.219.250.8:62920 | ✓ 719ms | 否 | ✓ 1531ms | ✓ 1715ms | 否 | http |
| 217.76.245.80:999 | ✓ 701ms | 否 | ✓ 1101ms | ✓ 1249ms | ✓ 1095ms | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1937ms | ✓ 1291ms | ✓ 985ms | http |
| 35.225.22.61:80 | ✓ 420ms | 否 | ✓ 337ms | ✓ 1160ms | ✓ 951ms | http |
| 166.0.192.117:8888 | ✓ 403ms | 否 | ✓ 1274ms | 否 | ✓ 1255ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1938ms | ✓ 792ms | 否 | ✓ 814ms | http |
| 172.212.68.37:3128 | 否 | ✓ 1353ms | ✓ 491ms | 否 | ✓ 954ms | http |
| 62.113.119.14:8080 | ✓ 814ms | 否 | ✓ 937ms | 否 | ✓ 1577ms | http |
| 147.45.251.242:8888 | ✓ 929ms | 否 | ✓ 1418ms | 否 | ✓ 1886ms | http |
| 185.115.74.185:8080 | ✓ 733ms | ✓ 1805ms | ✓ 1324ms | 否 | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 935ms | ✓ 1255ms | ✓ 1227ms | http |
| 160.238.65.6:3128 | ✓ 412ms | 否 | ✓ 665ms | 否 | ✓ 1554ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1263ms | ✓ 1460ms | ✓ 1170ms | http |
| 81.70.169.194:80 | ✓ 1847ms | ✓ 1775ms | ✓ 1527ms | ✓ 1899ms | 否 | http |
| 45.88.0.113:3128 | ✓ 1625ms | 否 | 否 | ✓ 1541ms | ✓ 1606ms | http |
| 120.198.141.75:22222 | 否 | ✓ 1623ms | ✓ 1142ms | 否 | ✓ 1081ms | http |
| 144.31.25.69:21064 | ✓ 481ms | 否 | ✓ 831ms | 否 | ✓ 1954ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1732ms | 否 | ✓ 1455ms | ✓ 1920ms | http |
| 120.198.141.80:22222 | ✓ 1211ms | ✓ 1657ms | ✓ 1349ms | ✓ 1507ms | ✓ 1211ms | http |
| 120.240.35.177:22222 | ✓ 1253ms | ✓ 1748ms | ✓ 1267ms | 否 | 否 | http |
| 45.88.0.115:3128 | ✓ 499ms | 否 | ✓ 458ms | 否 | ✓ 948ms | http |
| 45.88.0.98:3128 | ✓ 516ms | 否 | ✓ 462ms | ✓ 1252ms | 否 | http |
| 160.238.65.2:3128 | ✓ 1668ms | 否 | ✓ 983ms | ✓ 1735ms | 否 | http |
| 121.128.121.54:3128 | ✓ 829ms | ✓ 1970ms | ✓ 1170ms | 否 | 否 | http |
| 45.140.147.82:1082 | ✓ 925ms | ✓ 1225ms | ✓ 1230ms | 否 | ✓ 1143ms | http |
| 45.136.198.40:3128 | ✓ 954ms | ✓ 1749ms | ✓ 1746ms | 否 | 否 | http |
| 14.56.177.44:3128 | 否 | ✓ 1357ms | ✓ 1392ms | 否 | ✓ 1957ms | http |
| 34.101.184.164:3128 | ✓ 1046ms | 否 | ✓ 1424ms | ✓ 1573ms | ✓ 1537ms | http |
| 38.180.2.107:3128 | ✓ 910ms | ✓ 1763ms | ✓ 1336ms | ✓ 1976ms | ✓ 1638ms | http |
| 162.240.154.26:3128 | ✓ 866ms | 否 | ✓ 510ms | ✓ 1594ms | 否 | http |
| 103.215.36.88:15247 | ✓ 1217ms | ✓ 1701ms | ✓ 1627ms | 否 | 否 | http |
| 5.75.196.26:40000 | ✓ 390ms | ✓ 1580ms | ✓ 1459ms | 否 | ✓ 1717ms | http |
| 103.166.185.54:3128 | ✓ 1975ms | 否 | ✓ 1400ms | 否 | ✓ 1126ms | http |
| 91.193.240.157:9877 | ✓ 947ms | 否 | ✓ 987ms | 否 | ✓ 1352ms | http |
| 120.238.159.230:22222 | ✓ 1160ms | ✓ 1446ms | 否 | 否 | ✓ 1192ms | http |
| 222.184.48.242:22222 | 否 | ✓ 1393ms | ✓ 1772ms | ✓ 1327ms | ✓ 1047ms | http |
| 45.88.0.116:3128 | ✓ 1525ms | 否 | ✓ 1604ms | ✓ 1968ms | 否 | http |
| 47.94.228.56:8090 | ✓ 1421ms | ✓ 1835ms | ✓ 1713ms | 否 | ✓ 1457ms | http |
| 101.32.244.83:8080 | ✓ 1598ms | 否 | ✓ 1172ms | ✓ 1523ms | ✓ 1472ms | http |
| 121.40.231.103:7890 | ✓ 1141ms | ✓ 1397ms | ✓ 1363ms | 否 | 否 | http |
| 121.43.196.213:8222 | ✓ 1085ms | ✓ 1293ms | ✓ 1035ms | ✓ 1334ms | ✓ 1045ms | http |
| 121.43.196.210:8222 | ✓ 1104ms | ✓ 1291ms | ✓ 1030ms | ✓ 1380ms | ✓ 1027ms | http |
| 114.55.226.123:10086 | ✓ 1201ms | 否 | ✓ 1178ms | ✓ 1439ms | ✓ 1191ms | http |
| 160.238.65.8:3128 | ✓ 1054ms | ✓ 1349ms | 否 | ✓ 1539ms | 否 | http |
| 160.238.65.4:3128 | ✓ 1784ms | 否 | ✓ 1620ms | 否 | ✓ 1551ms | http |
| 181.78.194.249:999 | ✓ 1242ms | ✓ 1885ms | ✓ 1601ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1152ms | 否 | ✓ 984ms | ✓ 1368ms | 否 | http |
| 115.76.5.32:10006 | 否 | 否 | ✓ 1743ms | ✓ 1973ms | ✓ 1841ms | http |
| 115.231.181.40:8128 | ✓ 1105ms | 否 | ✓ 1425ms | 否 | ✓ 1145ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1729ms | ✓ 1799ms | ✓ 1410ms | http |
| 113.255.59.226:8080 | ✓ 1773ms | 否 | ✓ 1650ms | ✓ 1366ms | ✓ 1459ms | http |
| 35.234.17.221:8080 | ✓ 997ms | ✓ 1558ms | 否 | 否 | ✓ 1048ms | http |
| 160.238.65.9:3128 | ✓ 537ms | 否 | ✓ 1130ms | ✓ 1524ms | 否 | http |
| 120.232.242.119:22222 | 否 | ✓ 1385ms | ✓ 1064ms | ✓ 1317ms | 否 | http |
| 103.82.23.118:5242 | ✓ 1701ms | 否 | ✓ 1516ms | 否 | ✓ 1662ms | http |
| 113.59.32.162:22222 | 否 | ✓ 1881ms | ✓ 1575ms | 否 | ✓ 1560ms | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1539ms | ✓ 1433ms | ✓ 1718ms | http |
| 20.120.225.109:3128 | ✓ 1830ms | 否 | ✓ 1906ms | 否 | ✓ 1894ms | http |
| 94.158.49.82:3128 | 否 | ✓ 1982ms | ✓ 1559ms | 否 | ✓ 1785ms | http |
| 95.85.252.153:21064 | 否 | 否 | ✓ 1656ms | ✓ 1924ms | ✓ 1286ms | http |
| 120.198.141.79:22222 | 否 | 否 | ✓ 1323ms | ✓ 1475ms | ✓ 1143ms | http |
| 195.123.209.48:3128 | ✓ 769ms | 否 | ✓ 1706ms | ✓ 1896ms | ✓ 1723ms | http |
| 103.177.199.101:1111 | 否 | 否 | ✓ 1957ms | ✓ 1793ms | ✓ 1601ms | http |
| 45.177.178.23:999 | ✓ 869ms | 否 | ✓ 366ms | 否 | ✓ 994ms | http |
| 192.71.213.85:9091 | ✓ 656ms | 否 | ✓ 497ms | ✓ 1701ms | 否 | http |
| 192.71.213.85:9812 | ✓ 1178ms | 否 | ✓ 392ms | ✓ 1470ms | 否 | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1141ms | ✓ 1213ms | ✓ 1585ms | http |
| 183.249.5.117:22222 | 否 | ✓ 1576ms | 否 | ✓ 1612ms | ✓ 1494ms | http |
| 117.159.239.51:22222 | 否 | ✓ 1246ms | ✓ 1000ms | ✓ 1360ms | ✓ 1037ms | http |
| 200.69.83.203:999 | ✓ 1917ms | 否 | ✓ 1581ms | 否 | ✓ 1770ms | http |

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
