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

最后更新：2026-03-21 00:24:55 UTC（2026-03-21 08:24:55 UTC+8）

**代理总数：186**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 185 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 186 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.105:8449 | ✓ 520ms | ✓ 1013ms | ✓ 498ms | ✓ 989ms | ✓ 752ms | http |
| 38.34.179.63:8448 | ✓ 698ms | ✓ 994ms | ✓ 1130ms | ✓ 1028ms | ✓ 790ms | http |
| 43.99.54.236:5555 | ✓ 906ms | ✓ 1183ms | ✓ 844ms | ✓ 1069ms | ✓ 847ms | http |
| 147.161.210.140:8800 | ✓ 1592ms | ✓ 1418ms | ✓ 1056ms | ✓ 1075ms | ✓ 1163ms | http |
| 194.147.115.50:3128 | ✓ 1340ms | ✓ 1834ms | ✓ 1782ms | ✓ 1339ms | ✓ 1676ms | http |
| 174.138.24.77:1080 | ✓ 1411ms | 否 | ✓ 1630ms | ✓ 1275ms | ✓ 1011ms | http |
| 167.103.34.108:8800 | ✓ 1650ms | 否 | ✓ 1588ms | ✓ 1920ms | ✓ 1555ms | http |
| 113.160.132.26:8080 | ✓ 1671ms | 否 | 否 | ✓ 1480ms | ✓ 1244ms | http |
| 38.34.179.96:8451 | ✓ 1033ms | ✓ 994ms | ✓ 1281ms | ✓ 1004ms | ✓ 774ms | http |
| 38.95.77.16:6005 | ✓ 973ms | ✓ 1045ms | ✓ 1268ms | ✓ 1389ms | ✓ 1403ms | http |
| 219.117.204.211:7799 | ✓ 648ms | ✓ 1224ms | ✓ 1232ms | ✓ 1029ms | 否 | http |
| 137.220.150.22:6005 | ✓ 933ms | 否 | ✓ 924ms | ✓ 1384ms | ✓ 1086ms | http |
| 38.34.179.60:8450 | ✓ 953ms | 否 | ✓ 1714ms | ✓ 1194ms | 否 | http |
| 38.145.208.244:8448 | ✓ 1259ms | 否 | ✓ 1278ms | ✓ 1748ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1065ms | 否 | 否 | ✓ 1668ms | ✓ 1352ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1434ms | ✓ 1185ms | 否 | ✓ 1125ms | http |
| 38.145.218.228:8447 | ✓ 967ms | 否 | ✓ 708ms | 否 | ✓ 1643ms | http |
| 38.34.179.49:8450 | ✓ 1009ms | 否 | ✓ 344ms | ✓ 982ms | 否 | http |
| 38.34.179.150:8449 | ✓ 1002ms | ✓ 968ms | ✓ 1948ms | 否 | 否 | http |
| 103.183.10.169:3125 | 否 | 否 | ✓ 1916ms | ✓ 1912ms | ✓ 1621ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 268ms | ✓ 1109ms | ✓ 962ms | http |
| 45.136.131.39:8451 | ✓ 426ms | ✓ 1011ms | ✓ 363ms | ✓ 1017ms | ✓ 716ms | http |
| 115.231.181.40:8128 | ✓ 1285ms | ✓ 1373ms | ✓ 1180ms | 否 | ✓ 1103ms | http |
| 190.12.150.244:999 | ✓ 800ms | ✓ 1703ms | ✓ 991ms | ✓ 1589ms | ✓ 1350ms | http |
| 167.103.31.122:8800 | ✓ 1965ms | 否 | ✓ 1303ms | ✓ 1905ms | ✓ 1445ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1676ms | ✓ 1095ms | ✓ 1700ms | 否 | http |
| 147.161.239.240:8800 | 否 | 否 | ✓ 885ms | ✓ 1688ms | ✓ 1252ms | http |
| 45.136.131.35:8452 | ✓ 1299ms | ✓ 1087ms | ✓ 645ms | ✓ 1299ms | ✓ 1112ms | http |
| 101.43.127.100:8877 | ✓ 1022ms | ✓ 1387ms | ✓ 1678ms | ✓ 1384ms | ✓ 1116ms | http |
| 91.238.105.64:2024 | ✓ 1094ms | 否 | ✓ 1528ms | 否 | ✓ 1771ms | http |
| 217.174.244.117:3129 | ✓ 1418ms | 否 | ✓ 1999ms | 否 | ✓ 1963ms | http |
| 38.34.179.162:8451 | ✓ 596ms | ✓ 1536ms | ✓ 354ms | ✓ 953ms | ✓ 760ms | http |
| 38.34.179.14:8450 | ✓ 436ms | ✓ 1463ms | ✓ 1244ms | ✓ 987ms | ✓ 800ms | http |
| 38.34.179.20:8445 | ✓ 1844ms | ✓ 1419ms | ✓ 288ms | ✓ 974ms | ✓ 866ms | http |
| 38.145.220.33:8448 | ✓ 1600ms | 否 | ✓ 1368ms | ✓ 1502ms | ✓ 1891ms | http |
| 38.145.208.165:8450 | ✓ 489ms | ✓ 1115ms | ✓ 340ms | ✓ 1134ms | ✓ 764ms | http |
| 38.34.179.95:8450 | ✓ 543ms | ✓ 1472ms | ✓ 346ms | ✓ 983ms | 否 | http |
| 38.145.208.239:8445 | ✓ 379ms | ✓ 1470ms | ✓ 326ms | ✓ 1460ms | ✓ 756ms | http |
| 45.136.130.246:8445 | ✓ 345ms | 否 | ✓ 735ms | ✓ 1328ms | 否 | http |
| 38.145.208.215:8451 | ✓ 738ms | ✓ 1362ms | ✓ 1199ms | ✓ 1592ms | 否 | http |
| 137.184.1.87:3128 | ✓ 496ms | ✓ 1819ms | ✓ 1127ms | ✓ 986ms | ✓ 752ms | http |
| 45.136.130.186:8451 | ✓ 762ms | ✓ 1391ms | ✓ 898ms | ✓ 1164ms | ✓ 953ms | http |
| 38.34.183.224:8448 | ✓ 1011ms | ✓ 904ms | ✓ 400ms | ✓ 1131ms | ✓ 1111ms | http |
| 133.242.138.34:8100 | ✓ 1505ms | ✓ 1326ms | ✓ 901ms | ✓ 1197ms | ✓ 912ms | http |
| 137.220.150.170:6005 | ✓ 1427ms | 否 | ✓ 1226ms | ✓ 1397ms | ✓ 1601ms | http |
| 38.34.179.229:8445 | ✓ 1971ms | ✓ 1572ms | 否 | ✓ 1817ms | 否 | http |
| 38.34.183.222:8453 | ✓ 1246ms | ✓ 923ms | ✓ 842ms | 否 | ✓ 1359ms | http |
| 45.136.130.177:8448 | ✓ 1195ms | ✓ 1025ms | ✓ 815ms | 否 | 否 | http |
| 45.136.130.182:8443 | ✓ 1468ms | ✓ 1117ms | ✓ 878ms | 否 | 否 | http |
| 45.136.130.179:8443 | ✓ 1460ms | ✓ 1130ms | ✓ 897ms | 否 | 否 | http |
| 45.136.130.183:8443 | ✓ 1458ms | ✓ 1133ms | ✓ 895ms | 否 | 否 | http |
| 114.231.73.92:1080 | ✓ 1311ms | ✓ 1624ms | ✓ 1163ms | ✓ 1641ms | ✓ 1150ms | http |
| 38.34.183.225:8450 | ✓ 765ms | ✓ 950ms | ✓ 743ms | ✓ 1017ms | ✓ 706ms | http |
| 38.145.220.11:8445 | ✓ 782ms | ✓ 1057ms | ✓ 916ms | ✓ 1069ms | ✓ 795ms | http |
| 20.78.213.56:80 | ✓ 943ms | ✓ 1307ms | 否 | ✓ 1632ms | ✓ 1093ms | http |
| 38.34.179.83:8448 | 否 | ✓ 1230ms | 否 | ✓ 1769ms | ✓ 1293ms | http |
| 38.34.179.61:8445 | ✓ 1778ms | ✓ 1542ms | ✓ 1496ms | 否 | 否 | http |
| 38.145.218.9:8443 | ✓ 486ms | ✓ 947ms | ✓ 966ms | ✓ 845ms | ✓ 737ms | http |
| 38.145.208.243:8443 | ✓ 474ms | ✓ 965ms | ✓ 961ms | ✓ 949ms | ✓ 681ms | http |
| 38.145.208.246:8443 | ✓ 503ms | ✓ 989ms | ✓ 910ms | ✓ 934ms | ✓ 719ms | http |
| 38.145.220.39:8453 | ✓ 960ms | ✓ 1947ms | ✓ 456ms | ✓ 1008ms | ✓ 1513ms | http |
| 38.34.179.176:8443 | ✓ 788ms | ✓ 1035ms | ✓ 318ms | ✓ 975ms | ✓ 742ms | http |
| 38.34.179.177:8447 | ✓ 797ms | ✓ 1012ms | ✓ 311ms | ✓ 955ms | ✓ 708ms | http |
| 38.34.179.174:8443 | ✓ 786ms | ✓ 989ms | ✓ 338ms | ✓ 919ms | ✓ 756ms | http |
| 45.136.131.36:8450 | ✓ 796ms | ✓ 885ms | ✓ 1292ms | ✓ 1226ms | ✓ 748ms | http |
| 38.34.179.79:8450 | ✓ 1838ms | ✓ 928ms | ✓ 304ms | ✓ 1096ms | 否 | http |
| 137.220.151.110:6005 | ✓ 1607ms | 否 | ✓ 1094ms | ✓ 1319ms | ✓ 1706ms | http |
| 45.136.130.185:8443 | ✓ 307ms | ✓ 847ms | ✓ 300ms | ✓ 907ms | ✓ 720ms | http |
| 162.240.154.26:3128 | ✓ 1065ms | 否 | ✓ 1464ms | ✓ 1307ms | 否 | http |
| 38.34.178.152:8447 | ✓ 451ms | ✓ 997ms | ✓ 375ms | ✓ 1117ms | ✓ 840ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1434ms | ✓ 1920ms | ✓ 1783ms | ✓ 1236ms | http |
| 38.34.183.233:8448 | ✓ 1881ms | ✓ 1529ms | 否 | ✓ 1293ms | ✓ 729ms | http |
| 45.136.131.39:8443 | ✓ 372ms | ✓ 1318ms | ✓ 904ms | ✓ 1031ms | ✓ 1411ms | http |
| 38.34.179.26:8450 | ✓ 362ms | ✓ 1357ms | 否 | ✓ 1152ms | 否 | http |
| 45.136.130.172:8443 | ✓ 1231ms | ✓ 1069ms | ✓ 1162ms | 否 | ✓ 718ms | http |
| 38.34.179.16:8451 | ✓ 1545ms | ✓ 872ms | ✓ 481ms | ✓ 1105ms | 否 | http |
| 158.160.215.167:8124 | ✓ 974ms | 否 | ✓ 1263ms | ✓ 1874ms | ✓ 1727ms | http |
| 38.34.183.13:8449 | ✓ 304ms | ✓ 895ms | ✓ 860ms | 否 | 否 | http |
| 181.78.194.249:999 | ✓ 1193ms | ✓ 1901ms | ✓ 1406ms | ✓ 1750ms | ✓ 1781ms | http |
| 38.145.208.193:8443 | ✓ 955ms | ✓ 885ms | ✓ 397ms | ✓ 909ms | ✓ 693ms | http |
| 45.136.130.173:8447 | ✓ 929ms | ✓ 934ms | ✓ 348ms | ✓ 947ms | ✓ 738ms | http |
| 45.136.130.174:8447 | ✓ 925ms | ✓ 907ms | ✓ 376ms | ✓ 886ms | ✓ 743ms | http |
| 45.136.130.176:8447 | ✓ 924ms | ✓ 945ms | ✓ 337ms | ✓ 955ms | ✓ 709ms | http |
| 45.136.130.181:8443 | ✓ 920ms | ✓ 898ms | ✓ 381ms | ✓ 958ms | ✓ 690ms | http |
| 45.136.130.180:8443 | ✓ 921ms | ✓ 901ms | ✓ 381ms | ✓ 968ms | ✓ 736ms | http |
| 45.136.130.178:8443 | ✓ 925ms | ✓ 883ms | ✓ 399ms | ✓ 964ms | ✓ 722ms | http |
| 38.145.208.195:8443 | ✓ 955ms | ✓ 890ms | ✓ 392ms | ✓ 1090ms | ✓ 796ms | http |
| 38.145.208.192:8443 | ✓ 959ms | ✓ 884ms | ✓ 395ms | ✓ 1078ms | ✓ 814ms | http |
| 38.145.220.198:8448 | ✓ 953ms | ✓ 902ms | ✓ 360ms | ✓ 1129ms | ✓ 739ms | http |
| 38.34.183.234:8450 | ✓ 985ms | ✓ 906ms | ✓ 551ms | ✓ 950ms | ✓ 761ms | http |
| 45.136.130.187:8443 | ✓ 917ms | ✓ 1469ms | ✓ 319ms | ✓ 966ms | ✓ 732ms | http |
| 38.145.203.41:8443 | ✓ 1757ms | ✓ 1449ms | ✓ 265ms | ✓ 959ms | ✓ 749ms | http |
| 38.145.208.169:8443 | 否 | ✓ 1573ms | ✓ 675ms | ✓ 919ms | ✓ 723ms | http |
| 38.145.208.172:8443 | 否 | ✓ 1522ms | ✓ 702ms | ✓ 960ms | ✓ 734ms | http |
| 38.145.208.173:8443 | 否 | ✓ 1482ms | ✓ 659ms | ✓ 946ms | ✓ 796ms | http |
| 38.145.208.177:8443 | 否 | ✓ 1497ms | ✓ 702ms | ✓ 939ms | ✓ 748ms | http |
| 38.145.208.170:8443 | 否 | ✓ 1547ms | ✓ 700ms | ✓ 966ms | ✓ 722ms | http |
| 1.225.116.115:1080 | ✓ 1709ms | 否 | ✓ 1000ms | ✓ 1093ms | ✓ 896ms | http |
| 38.34.179.76:8450 | 否 | ✓ 1064ms | ✓ 346ms | ✓ 1499ms | ✓ 1603ms | http |
| 217.160.224.54:8118 | 否 | ✓ 1770ms | ✓ 1809ms | 否 | ✓ 1723ms | http |

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
