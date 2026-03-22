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

最后更新：2026-03-22 07:52:10 UTC（2026-03-22 15:52:10 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 436ms | 否 | ✓ 433ms | ✓ 1164ms | ✓ 1173ms | http |
| 43.99.54.236:5555 | ✓ 629ms | 否 | ✓ 611ms | ✓ 766ms | ✓ 644ms | http |
| 147.161.210.140:8800 | 否 | ✓ 1600ms | ✓ 839ms | ✓ 956ms | ✓ 943ms | http |
| 162.240.154.26:3128 | ✓ 942ms | ✓ 1898ms | 否 | ✓ 1362ms | ✓ 1099ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1069ms | ✓ 1697ms | ✓ 1145ms | ✓ 929ms | http |
| 8.219.97.248:80 | ✓ 1430ms | 否 | ✓ 1470ms | ✓ 1625ms | ✓ 1158ms | http |
| 113.160.132.26:8080 | ✓ 1724ms | 否 | ✓ 1298ms | ✓ 1184ms | ✓ 986ms | http |
| 147.161.239.240:8800 | ✓ 1359ms | 否 | ✓ 1616ms | ✓ 1815ms | ✓ 1566ms | http |
| 106.75.15.167:7890 | ✓ 1100ms | ✓ 1361ms | 否 | ✓ 1346ms | 否 | http |
| 183.249.5.117:22222 | ✓ 688ms | ✓ 887ms | ✓ 860ms | ✓ 1072ms | ✓ 888ms | http |
| 34.96.238.40:8080 | ✓ 1287ms | 否 | ✓ 1012ms | 否 | ✓ 1107ms | http |
| 167.103.34.108:8800 | ✓ 1077ms | 否 | ✓ 1129ms | ✓ 1349ms | ✓ 1201ms | http |
| 137.220.150.22:6005 | ✓ 883ms | 否 | ✓ 845ms | ✓ 1301ms | 否 | http |
| 219.117.204.211:7799 | ✓ 1153ms | 否 | ✓ 606ms | ✓ 934ms | ✓ 1386ms | http |
| 45.151.183.183:1080 | ✓ 925ms | 否 | ✓ 1478ms | ✓ 1814ms | ✓ 1489ms | http |
| 77.232.135.22:1080 | ✓ 842ms | 否 | ✓ 1589ms | ✓ 1595ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1358ms | 否 | ✓ 1176ms | ✓ 940ms | http |
| 222.184.48.251:22222 | ✓ 856ms | ✓ 1129ms | ✓ 931ms | ✓ 1919ms | 否 | http |
| 142.171.224.229:7890 | ✓ 897ms | ✓ 764ms | ✓ 370ms | ✓ 739ms | ✓ 490ms | http |
| 45.140.147.155:1082 | ✓ 630ms | 否 | ✓ 1504ms | ✓ 1913ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1928ms | 否 | ✓ 1420ms | ✓ 1971ms | ✓ 1501ms | http |
| 91.238.105.64:2024 | ✓ 801ms | 否 | ✓ 1104ms | ✓ 1781ms | ✓ 1591ms | http |
| 103.84.95.54:7890 | ✓ 615ms | 否 | 否 | ✓ 814ms | ✓ 623ms | http |
| 38.145.218.87:8444 | ✓ 888ms | ✓ 1326ms | ✓ 757ms | ✓ 684ms | ✓ 771ms | http |
| 103.191.92.157:1009 | ✓ 1611ms | 否 | ✓ 1291ms | ✓ 1256ms | ✓ 913ms | http |
| 160.250.5.22:1 | ✓ 1632ms | 否 | ✓ 1369ms | ✓ 1396ms | ✓ 1039ms | http |
| 45.136.198.40:3128 | ✓ 1337ms | 否 | ✓ 1688ms | 否 | ✓ 1816ms | http |
| 101.43.127.100:8877 | ✓ 813ms | ✓ 1028ms | 否 | 否 | ✓ 1762ms | http |
| 104.168.158.236:10808 | ✓ 853ms | 否 | ✓ 1680ms | ✓ 1603ms | 否 | http |
| 183.249.5.111:22222 | ✓ 838ms | ✓ 1117ms | ✓ 1060ms | ✓ 1372ms | ✓ 723ms | http |
| 160.250.4.245:1 | ✓ 1611ms | 否 | ✓ 1624ms | ✓ 1496ms | ✓ 1124ms | http |
| 103.39.51.207:8080 | ✓ 1927ms | 否 | ✓ 1210ms | ✓ 1560ms | ✓ 1489ms | http |
| 137.220.150.104:6005 | ✓ 752ms | 否 | ✓ 947ms | ✓ 1082ms | ✓ 907ms | http |
| 72.56.79.129:1080 | ✓ 1420ms | 否 | 否 | ✓ 1971ms | ✓ 1653ms | http |
| 185.241.5.57:3128 | ✓ 1165ms | 否 | ✓ 1098ms | 否 | ✓ 1711ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1121ms | ✓ 1218ms | ✓ 959ms | http |
| 183.249.5.105:22222 | ✓ 1233ms | ✓ 1160ms | ✓ 1020ms | ✓ 1145ms | ✓ 981ms | http |
| 45.168.238.193:8443 | ✓ 667ms | 否 | ✓ 621ms | ✓ 1301ms | ✓ 1039ms | http |
| 45.140.147.155:1081 | ✓ 1209ms | 否 | ✓ 1218ms | 否 | ✓ 1381ms | http |
| 120.92.212.16:7890 | ✓ 1952ms | ✓ 1168ms | 否 | ✓ 1841ms | 否 | http |
| 217.76.245.80:999 | ✓ 1081ms | ✓ 1534ms | ✓ 1357ms | ✓ 1617ms | ✓ 1491ms | http |
| 101.32.244.83:8080 | ✓ 944ms | 否 | ✓ 881ms | ✓ 1354ms | ✓ 1152ms | http |
| 121.43.196.210:8222 | ✓ 936ms | ✓ 1059ms | ✓ 802ms | ✓ 1087ms | ✓ 880ms | http |
| 121.43.196.213:8222 | ✓ 951ms | ✓ 1047ms | ✓ 937ms | ✓ 1055ms | ✓ 841ms | http |
| 114.55.226.123:10086 | ✓ 1046ms | ✓ 1629ms | ✓ 969ms | ✓ 1182ms | ✓ 1022ms | http |
| 38.34.179.189:8452 | 否 | ✓ 1765ms | ✓ 1362ms | ✓ 1989ms | ✓ 994ms | http |
| 183.249.5.110:22222 | ✓ 667ms | ✓ 858ms | ✓ 750ms | ✓ 933ms | ✓ 981ms | http |
| 144.31.79.117:8888 | ✓ 1130ms | 否 | ✓ 1391ms | ✓ 1974ms | ✓ 1553ms | http |
| 166.88.55.83:7890 | ✓ 612ms | ✓ 1057ms | ✓ 598ms | ✓ 753ms | ✓ 604ms | http |
| 34.150.20.6:8888 | ✓ 1003ms | ✓ 1443ms | ✓ 938ms | ✓ 785ms | ✓ 620ms | http |
| 116.80.49.167:3172 | ✓ 1505ms | 否 | ✓ 1508ms | ✓ 1767ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1230ms | 否 | ✓ 1742ms | 否 | ✓ 1890ms | http |
| 116.80.49.159:3172 | 否 | 否 | ✓ 1497ms | ✓ 1775ms | ✓ 1620ms | http |
| 137.220.150.170:6005 | ✓ 1394ms | 否 | ✓ 971ms | ✓ 1284ms | ✓ 1003ms | http |
| 45.186.6.104:3128 | ✓ 1380ms | ✓ 1689ms | ✓ 1897ms | 否 | 否 | http |
| 38.145.208.210:8448 | ✓ 1453ms | 否 | ✓ 863ms | ✓ 1528ms | ✓ 1825ms | http |
| 172.212.68.37:3128 | ✓ 678ms | 否 | ✓ 284ms | ✓ 1829ms | ✓ 1636ms | http |
| 85.208.108.43:2094 | ✓ 1009ms | 否 | ✓ 962ms | ✓ 1637ms | ✓ 1027ms | http |
| 181.78.44.63:999 | 否 | 否 | ✓ 1581ms | ✓ 1482ms | ✓ 1258ms | http |
| 103.39.51.190:8080 | ✓ 1651ms | 否 | 否 | ✓ 1422ms | ✓ 1825ms | http |
| 103.82.23.118:5242 | ✓ 1207ms | ✓ 1838ms | ✓ 1306ms | 否 | ✓ 1613ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1513ms | ✓ 760ms | ✓ 1835ms | 否 | http |
| 150.241.77.172:1080 | ✓ 700ms | 否 | ✓ 1492ms | ✓ 1874ms | ✓ 1332ms | http |
| 47.77.193.180:1080 | ✓ 112ms | ✓ 881ms | ✓ 379ms | ✓ 654ms | ✓ 498ms | http |
| 43.165.195.107:3128 | ✓ 841ms | 否 | ✓ 863ms | ✓ 1133ms | ✓ 933ms | http |
| 194.67.99.223:1080 | 否 | 否 | ✓ 1913ms | ✓ 1954ms | ✓ 1700ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1187ms | ✓ 870ms | 否 | ✓ 871ms | http |
| 38.145.203.97:8444 | ✓ 1605ms | ✓ 1823ms | ✓ 441ms | ✓ 914ms | ✓ 1315ms | http |
| 180.103.19.207:1080 | ✓ 962ms | ✓ 1644ms | ✓ 1507ms | 否 | ✓ 1230ms | http |
| 222.184.48.252:22222 | 否 | 否 | ✓ 952ms | ✓ 1222ms | ✓ 877ms | http |
| 38.145.208.182:8446 | ✓ 952ms | 否 | ✓ 357ms | ✓ 1287ms | 否 | http |
| 121.230.8.213:1080 | ✓ 1054ms | 否 | 否 | ✓ 1381ms | ✓ 1012ms | http |
| 121.230.9.241:1080 | ✓ 1319ms | 否 | ✓ 1192ms | ✓ 1646ms | 否 | http |
| 103.144.102.231:8085 | ✓ 1672ms | 否 | ✓ 1913ms | 否 | ✓ 1696ms | http |
| 45.136.131.55:8444 | ✓ 492ms | ✓ 715ms | 否 | ✓ 1527ms | ✓ 637ms | http |
| 45.136.130.167:8446 | ✓ 492ms | ✓ 1172ms | ✓ 1146ms | ✓ 1998ms | ✓ 520ms | http |
| 38.34.178.155:8448 | 否 | ✓ 1431ms | ✓ 990ms | 否 | ✓ 522ms | http |
| 202.141.161.53:30001 | ✓ 993ms | ✓ 1281ms | ✓ 1673ms | 否 | ✓ 1153ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1852ms | ✓ 1965ms | ✓ 1269ms | ✓ 952ms | http |
| 121.230.9.252:1080 | ✓ 1364ms | ✓ 1981ms | ✓ 1123ms | ✓ 1970ms | 否 | http |

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
