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

最后更新：2026-03-29 03:46:07 UTC（2026-03-29 11:46:07 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 655ms | ✓ 975ms | ✓ 709ms | ✓ 776ms | ✓ 620ms | http |
| 39.185.46.193:5911 | ✓ 720ms | ✓ 812ms | ✓ 680ms | ✓ 858ms | ✓ 649ms | http |
| 35.225.22.61:80 | ✓ 618ms | ✓ 1407ms | ✓ 570ms | ✓ 1174ms | ✓ 1185ms | http |
| 147.161.210.140:8800 | ✓ 1742ms | ✓ 1045ms | ✓ 742ms | ✓ 1132ms | ✓ 938ms | http |
| 115.231.181.40:8128 | ✓ 810ms | ✓ 1811ms | ✓ 945ms | ✓ 1045ms | ✓ 893ms | http |
| 113.160.132.26:8080 | ✓ 1804ms | ✓ 1290ms | 否 | ✓ 1400ms | ✓ 968ms | http |
| 167.103.115.102:8800 | ✓ 1579ms | 否 | ✓ 1207ms | ✓ 1245ms | ✓ 1931ms | http |
| 167.103.34.108:8800 | ✓ 1605ms | 否 | ✓ 1476ms | ✓ 1417ms | ✓ 1652ms | http |
| 95.213.217.168:52004 | ✓ 1343ms | 否 | ✓ 1667ms | 否 | ✓ 1820ms | http |
| 103.9.78.2:3128 | ✓ 1592ms | ✓ 1819ms | ✓ 1688ms | 否 | ✓ 1654ms | http |
| 208.87.243.199:7878 | ✓ 919ms | ✓ 680ms | ✓ 654ms | ✓ 703ms | ✓ 713ms | http |
| 5.104.87.17:8051 | ✓ 699ms | 否 | ✓ 755ms | ✓ 887ms | ✓ 710ms | http |
| 103.84.95.54:7890 | ✓ 1567ms | 否 | ✓ 648ms | 否 | ✓ 999ms | http |
| 167.103.144.127:8800 | ✓ 1513ms | 否 | ✓ 1368ms | ✓ 1742ms | ✓ 1489ms | http |
| 34.101.184.164:3128 | ✓ 1241ms | 否 | ✓ 1519ms | 否 | ✓ 1966ms | http |
| 45.12.151.226:2829 | ✓ 925ms | ✓ 1743ms | 否 | 否 | ✓ 1904ms | http |
| 106.75.15.167:7890 | ✓ 1134ms | ✓ 1362ms | ✓ 1711ms | ✓ 1502ms | 否 | http |
| 59.46.216.131:30001 | ✓ 919ms | ✓ 1272ms | 否 | 否 | ✓ 988ms | http |
| 180.250.219.58:53281 | ✓ 1678ms | 否 | ✓ 1352ms | ✓ 1807ms | ✓ 1751ms | http |
| 167.103.31.122:8800 | ✓ 1429ms | 否 | ✓ 1293ms | ✓ 1611ms | ✓ 1522ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1064ms | ✓ 806ms | ✓ 640ms | http |
| 42.96.16.158:1311 | ✓ 1009ms | 否 | ✓ 1135ms | ✓ 1187ms | 否 | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 875ms | ✓ 1255ms | ✓ 827ms | http |
| 101.43.127.100:8877 | ✓ 834ms | ✓ 1063ms | ✓ 827ms | ✓ 1090ms | ✓ 848ms | http |
| 128.199.116.219:9090 | ✓ 1158ms | 否 | ✓ 772ms | ✓ 1048ms | ✓ 788ms | http |
| 128.199.113.85:9090 | ✓ 1162ms | 否 | ✓ 962ms | ✓ 1023ms | ✓ 822ms | http |
| 147.161.239.240:8800 | ✓ 1419ms | ✓ 1855ms | ✓ 1436ms | ✓ 1855ms | ✓ 1632ms | http |
| 128.199.114.189:9090 | ✓ 1568ms | 否 | ✓ 1288ms | ✓ 1747ms | ✓ 964ms | http |
| 120.92.212.16:8890 | ✓ 976ms | ✓ 1158ms | ✓ 930ms | 否 | ✓ 960ms | http |
| 177.234.217.88:999 | 否 | 否 | ✓ 1711ms | ✓ 1967ms | ✓ 1744ms | http |
| 38.145.218.87:8445 | 否 | ✓ 1098ms | ✓ 257ms | ✓ 1383ms | 否 | http |
| 128.199.254.13:9090 | ✓ 739ms | 否 | 否 | ✓ 1501ms | ✓ 841ms | http |
| 120.92.212.16:7890 | ✓ 981ms | ✓ 1161ms | ✓ 1188ms | ✓ 1391ms | ✓ 1723ms | http |
| 160.250.5.22:1 | ✓ 1336ms | 否 | ✓ 1237ms | ✓ 1176ms | ✓ 1030ms | http |
| 91.238.123.230:8000 | ✓ 1101ms | ✓ 1527ms | ✓ 1298ms | ✓ 1679ms | ✓ 1603ms | http |
| 1.231.81.166:3128 | ✓ 1421ms | ✓ 1840ms | ✓ 1557ms | ✓ 1488ms | ✓ 1374ms | http |
| 47.95.231.180:8084 | ✓ 829ms | ✓ 1142ms | ✓ 968ms | ✓ 1138ms | ✓ 1005ms | http |
| 45.8.157.38:3128 | ✓ 606ms | ✓ 1363ms | ✓ 1077ms | 否 | 否 | http |
| 193.233.22.29:10808 | ✓ 975ms | 否 | ✓ 1462ms | ✓ 1433ms | 否 | http |
| 5.102.109.41:999 | ✓ 678ms | ✓ 1406ms | ✓ 961ms | ✓ 1413ms | ✓ 1160ms | http |
| 38.145.208.169:8445 | ✓ 713ms | 否 | ✓ 355ms | ✓ 660ms | 否 | http |
| 38.145.208.174:8445 | ✓ 755ms | 否 | ✓ 309ms | ✓ 712ms | 否 | http |
| 38.145.218.229:8453 | ✓ 731ms | 否 | ✓ 736ms | ✓ 1183ms | 否 | http |
| 146.190.80.158:9090 | ✓ 1127ms | 否 | ✓ 1072ms | ✓ 1058ms | ✓ 894ms | http |
| 38.145.208.164:8451 | ✓ 1054ms | ✓ 893ms | ✓ 1146ms | 否 | ✓ 620ms | http |
| 103.203.234.105:3127 | 否 | 否 | ✓ 1200ms | ✓ 1340ms | ✓ 1322ms | http |
| 38.145.218.216:8452 | ✓ 649ms | ✓ 676ms | ✓ 962ms | 否 | ✓ 702ms | http |
| 38.34.179.67:8451 | ✓ 190ms | ✓ 1355ms | ✓ 152ms | ✓ 676ms | ✓ 866ms | http |
| 38.145.208.166:8451 | ✓ 1261ms | ✓ 690ms | ✓ 380ms | ✓ 1538ms | ✓ 1840ms | http |
| 38.145.220.100:8450 | ✓ 359ms | ✓ 611ms | ✓ 1322ms | 否 | ✓ 529ms | http |
| 38.145.208.204:8445 | ✓ 1677ms | ✓ 1474ms | ✓ 138ms | ✓ 652ms | ✓ 1401ms | http |
| 38.34.179.86:8452 | ✓ 825ms | 否 | ✓ 566ms | ✓ 856ms | ✓ 1822ms | http |
| 38.145.218.232:8445 | ✓ 296ms | ✓ 1439ms | ✓ 915ms | ✓ 660ms | ✓ 960ms | http |
| 38.145.218.227:8451 | ✓ 207ms | ✓ 1115ms | ✓ 1703ms | ✓ 685ms | ✓ 1168ms | http |
| 38.145.208.171:8451 | ✓ 230ms | ✓ 1258ms | ✓ 1569ms | ✓ 675ms | ✓ 1209ms | http |
| 38.34.179.186:8453 | ✓ 581ms | 否 | ✓ 793ms | ✓ 689ms | ✓ 583ms | http |
| 38.34.183.224:8451 | ✓ 183ms | ✓ 686ms | ✓ 1247ms | 否 | ✓ 499ms | http |
| 38.145.208.195:8453 | ✓ 492ms | ✓ 1572ms | ✓ 1573ms | ✓ 699ms | ✓ 565ms | http |
| 38.34.178.193:8446 | ✓ 1249ms | ✓ 881ms | ✓ 726ms | 否 | ✓ 1027ms | http |
| 38.34.179.12:8452 | ✓ 383ms | 否 | ✓ 1086ms | ✓ 1198ms | 否 | http |
| 38.145.208.171:8453 | ✓ 841ms | 否 | ✓ 877ms | 否 | ✓ 1951ms | http |
| 38.34.183.130:8452 | ✓ 255ms | ✓ 1075ms | ✓ 731ms | ✓ 1923ms | ✓ 670ms | http |
| 20.210.76.104:8561 | ✓ 1423ms | ✓ 1467ms | ✓ 1265ms | ✓ 1855ms | 否 | http |
| 20.27.15.49:8561 | ✓ 1397ms | ✓ 1470ms | ✓ 1262ms | ✓ 1855ms | 否 | http |
| 20.210.76.175:8561 | ✓ 1421ms | ✓ 1574ms | ✓ 1265ms | ✓ 1747ms | 否 | http |
| 116.80.49.167:3172 | ✓ 1843ms | 否 | ✓ 1481ms | ✓ 1782ms | 否 | http |
| 139.159.99.242:8080 | 否 | ✓ 943ms | ✓ 783ms | ✓ 1693ms | 否 | http |
| 118.31.1.154:80 | 否 | 否 | ✓ 1115ms | ✓ 1061ms | ✓ 851ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 571ms | ✓ 1087ms | ✓ 947ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 569ms | ✓ 1081ms | ✓ 954ms | http |
| 8.219.97.248:80 | ✓ 1621ms | 否 | ✓ 1877ms | 否 | ✓ 1063ms | http |
| 103.39.51.190:8080 | ✓ 1628ms | 否 | ✓ 1803ms | ✓ 1549ms | ✓ 1983ms | http |
| 103.234.27.194:8080 | ✓ 1830ms | 否 | 否 | ✓ 1905ms | ✓ 1787ms | http |
| 103.113.70.189:1081 | ✓ 482ms | ✓ 1395ms | ✓ 702ms | ✓ 1198ms | 否 | http |
| 202.40.177.46:8080 | ✓ 1824ms | 否 | ✓ 1849ms | ✓ 1615ms | ✓ 1548ms | http |
| 137.184.6.37:3128 | ✓ 230ms | ✓ 724ms | ✓ 755ms | ✓ 722ms | ✓ 473ms | http |
| 45.149.92.147:5001 | ✓ 587ms | 否 | ✓ 596ms | ✓ 769ms | ✓ 704ms | http |
| 116.80.49.165:3172 | 否 | 否 | ✓ 1472ms | ✓ 1787ms | ✓ 1625ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1239ms | ✓ 1685ms | ✓ 934ms | http |
| 121.230.8.45:1080 | ✓ 1532ms | ✓ 1450ms | 否 | ✓ 1388ms | ✓ 1326ms | http |
| 38.34.179.37:8451 | ✓ 1877ms | 否 | ✓ 1979ms | 否 | ✓ 1989ms | http |
| 104.247.51.76:3128 | ✓ 570ms | 否 | ✓ 912ms | ✓ 1283ms | ✓ 1181ms | http |
| 106.117.208.101:7890 | ✓ 948ms | ✓ 1225ms | 否 | ✓ 1814ms | ✓ 1127ms | http |
| 202.129.206.239:3128 | ✓ 1645ms | 否 | ✓ 1341ms | ✓ 1654ms | ✓ 1475ms | http |
| 38.34.179.105:8449 | ✓ 505ms | ✓ 706ms | ✓ 1557ms | 否 | ✓ 572ms | http |

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
