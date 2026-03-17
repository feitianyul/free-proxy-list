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

最后更新：2026-03-17 06:53:42 UTC（2026-03-17 14:53:42 UTC+8）

**代理总数：97**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 97 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 97 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 490ms | 否 | 否 | ✓ 1089ms | ✓ 1071ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1032ms | ✓ 1414ms | ✓ 1030ms | http |
| 62.60.177.204:34094 | ✓ 360ms | 否 | ✓ 1374ms | 否 | ✓ 1568ms | http |
| 147.161.210.140:8800 | 否 | ✓ 1344ms | ✓ 1086ms | ✓ 1316ms | ✓ 1014ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1361ms | 否 | ✓ 1137ms | ✓ 960ms | http |
| 45.167.124.52:8080 | ✓ 1095ms | ✓ 1930ms | ✓ 1606ms | 否 | ✓ 1483ms | http |
| 103.84.95.54:7890 | ✓ 813ms | 否 | ✓ 1095ms | 否 | ✓ 774ms | http |
| 47.101.149.27:9010 | ✓ 1520ms | 否 | 否 | ✓ 1641ms | ✓ 1428ms | http |
| 137.220.150.152:6005 | ✓ 1360ms | 否 | ✓ 989ms | ✓ 1306ms | ✓ 1029ms | http |
| 137.220.150.170:6005 | ✓ 1346ms | 否 | ✓ 1107ms | ✓ 1231ms | ✓ 1102ms | http |
| 168.235.110.63:3128 | 否 | ✓ 1303ms | ✓ 1065ms | ✓ 1267ms | ✓ 792ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1542ms | ✓ 1997ms | ✓ 1511ms | http |
| 52.74.26.202:8080 | ✓ 862ms | ✓ 1386ms | ✓ 855ms | ✓ 1259ms | ✓ 1113ms | http |
| 115.231.181.40:8128 | ✓ 1070ms | ✓ 1252ms | ✓ 1556ms | ✓ 1361ms | ✓ 1692ms | http |
| 45.136.198.40:3128 | ✓ 1108ms | ✓ 1947ms | ✓ 1776ms | 否 | ✓ 1677ms | http |
| 101.43.127.100:8877 | ✓ 1014ms | ✓ 1272ms | ✓ 931ms | ✓ 1892ms | ✓ 954ms | http |
| 113.160.132.26:8080 | ✓ 1116ms | ✓ 1821ms | 否 | ✓ 1516ms | ✓ 1090ms | http |
| 137.220.151.110:6005 | ✓ 1113ms | 否 | ✓ 1072ms | ✓ 1483ms | 否 | http |
| 85.198.96.242:3128 | ✓ 996ms | 否 | ✓ 1899ms | ✓ 1829ms | 否 | http |
| 160.238.65.4:3128 | ✓ 505ms | 否 | ✓ 531ms | ✓ 1384ms | ✓ 1111ms | http |
| 160.238.65.3:3128 | ✓ 899ms | ✓ 1355ms | ✓ 783ms | ✓ 1375ms | 否 | http |
| 160.238.65.2:3128 | ✓ 509ms | ✓ 1517ms | ✓ 1011ms | ✓ 1721ms | ✓ 1064ms | http |
| 160.238.65.5:3128 | ✓ 514ms | 否 | ✓ 848ms | ✓ 1357ms | ✓ 1056ms | http |
| 160.238.65.9:3128 | ✓ 513ms | 否 | ✓ 520ms | ✓ 1420ms | 否 | http |
| 72.11.151.159:6005 | ✓ 844ms | 否 | ✓ 1717ms | 否 | ✓ 1431ms | http |
| 165.227.5.10:8888 | ✓ 481ms | ✓ 1004ms | ✓ 928ms | 否 | 否 | http |
| 133.242.138.34:8100 | ✓ 867ms | ✓ 1291ms | ✓ 1873ms | 否 | 否 | http |
| 137.220.150.104:6005 | ✓ 1276ms | 否 | ✓ 1380ms | ✓ 1292ms | 否 | http |
| 45.88.0.115:3128 | ✓ 682ms | 否 | ✓ 508ms | ✓ 1349ms | ✓ 1228ms | http |
| 213.220.62.62:3128 | 否 | ✓ 1856ms | ✓ 1386ms | ✓ 1945ms | ✓ 1460ms | http |
| 45.88.0.114:3128 | ✓ 690ms | 否 | ✓ 510ms | ✓ 1356ms | ✓ 1213ms | http |
| 45.88.0.116:3128 | ✓ 661ms | 否 | ✓ 515ms | ✓ 1337ms | ✓ 1231ms | http |
| 212.192.12.90:6005 | 否 | 否 | ✓ 1202ms | ✓ 1221ms | ✓ 810ms | http |
| 45.88.0.113:3128 | ✓ 902ms | ✓ 1499ms | ✓ 1511ms | 否 | ✓ 1456ms | http |
| 45.88.0.111:3128 | ✓ 899ms | 否 | ✓ 1018ms | 否 | ✓ 1451ms | http |
| 45.88.0.99:3128 | ✓ 902ms | 否 | ✓ 1013ms | 否 | ✓ 1467ms | http |
| 45.88.0.98:3128 | ✓ 902ms | ✓ 1621ms | ✓ 1393ms | 否 | ✓ 1465ms | http |
| 45.88.0.117:3128 | ✓ 902ms | 否 | ✓ 1009ms | 否 | ✓ 1473ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1848ms | ✓ 1301ms | ✓ 1196ms | http |
| 38.145.218.113:8444 | ✓ 921ms | ✓ 1096ms | ✓ 576ms | ✓ 1184ms | ✓ 733ms | http |
| 38.145.208.205:8453 | ✓ 921ms | ✓ 958ms | ✓ 1316ms | ✓ 1073ms | ✓ 952ms | http |
| 45.136.130.191:8451 | ✓ 902ms | ✓ 1223ms | ✓ 1921ms | ✓ 1846ms | ✓ 708ms | http |
| 45.136.130.185:8445 | ✓ 900ms | ✓ 1116ms | ✓ 558ms | ✓ 1022ms | 否 | http |
| 45.136.131.43:8445 | ✓ 903ms | ✓ 1114ms | ✓ 1157ms | 否 | ✓ 772ms | http |
| 219.117.204.211:7799 | ✓ 1134ms | 否 | ✓ 1923ms | ✓ 1285ms | ✓ 747ms | http |
| 138.124.53.25:7443 | ✓ 1770ms | 否 | ✓ 1669ms | ✓ 1934ms | 否 | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1027ms | ✓ 1362ms | ✓ 1049ms | http |
| 120.92.212.16:7890 | ✓ 1006ms | ✓ 1306ms | ✓ 1036ms | ✓ 1293ms | ✓ 1050ms | http |
| 167.71.196.28:8080 | ✓ 1450ms | 否 | ✓ 1467ms | 否 | ✓ 1022ms | http |
| 1.225.116.115:1080 | ✓ 1953ms | 否 | ✓ 1357ms | ✓ 1358ms | ✓ 1019ms | http |
| 192.3.203.158:1080 | ✓ 1036ms | 否 | ✓ 1353ms | 否 | ✓ 1874ms | http |
| 59.46.216.131:30001 | ✓ 1169ms | 否 | ✓ 1149ms | 否 | ✓ 1565ms | http |
| 211.171.114.154:3128 | ✓ 1278ms | ✓ 1403ms | ✓ 1682ms | ✓ 1458ms | 否 | http |
| 172.212.68.37:3128 | ✓ 289ms | ✓ 1576ms | ✓ 891ms | ✓ 1379ms | ✓ 1218ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1375ms | ✓ 1426ms | ✓ 1196ms | http |
| 8.219.97.248:80 | ✓ 1633ms | 否 | ✓ 1477ms | 否 | ✓ 1351ms | http |
| 192.71.213.85:9091 | ✓ 719ms | 否 | ✓ 849ms | ✓ 1674ms | 否 | http |
| 112.163.160.93:3128 | 否 | ✓ 1831ms | ✓ 961ms | ✓ 1686ms | 否 | http |
| 47.77.193.180:1080 | ✓ 278ms | ✓ 1466ms | ✓ 351ms | ✓ 825ms | ✓ 639ms | http |
| 38.145.208.186:8444 | ✓ 613ms | ✓ 1175ms | ✓ 505ms | ✓ 965ms | ✓ 860ms | http |
| 38.145.208.184:8444 | ✓ 615ms | ✓ 1114ms | ✓ 563ms | ✓ 973ms | ✓ 858ms | http |
| 38.145.203.107:8453 | ✓ 613ms | ✓ 979ms | ✓ 700ms | ✓ 1097ms | ✓ 748ms | http |
| 38.145.208.217:8453 | ✓ 633ms | ✓ 1599ms | ✓ 359ms | ✓ 1022ms | ✓ 775ms | http |
| 38.145.203.34:8445 | ✓ 613ms | ✓ 1178ms | ✓ 501ms | ✓ 964ms | ✓ 762ms | http |
| 38.145.218.232:8445 | ✓ 614ms | ✓ 1521ms | ✓ 375ms | ✓ 966ms | ✓ 806ms | http |
| 38.145.203.46:8445 | ✓ 614ms | ✓ 1524ms | ✓ 385ms | ✓ 1021ms | ✓ 783ms | http |
| 38.145.208.177:8445 | ✓ 634ms | ✓ 1857ms | ✓ 365ms | ✓ 975ms | ✓ 763ms | http |
| 38.145.208.204:8453 | ✓ 633ms | ✓ 1733ms | ✓ 371ms | ✓ 1122ms | ✓ 812ms | http |
| 45.136.130.174:8452 | ✓ 611ms | ✓ 1534ms | ✓ 420ms | ✓ 1165ms | ✓ 953ms | http |
| 38.145.203.132:8445 | ✓ 1174ms | ✓ 1526ms | ✓ 372ms | ✓ 981ms | ✓ 735ms | http |
| 38.145.220.56:8450 | ✓ 612ms | 否 | ✓ 382ms | ✓ 1302ms | ✓ 769ms | http |
| 38.145.208.230:8445 | ✓ 1172ms | ✓ 1156ms | ✓ 316ms | 否 | ✓ 755ms | http |
| 38.145.218.234:8445 | ✓ 614ms | 否 | ✓ 378ms | ✓ 1972ms | ✓ 793ms | http |
| 38.145.208.172:8448 | ✓ 902ms | ✓ 1821ms | ✓ 483ms | ✓ 1285ms | ✓ 1555ms | http |
| 38.145.220.182:8445 | ✓ 614ms | ✓ 1086ms | ✓ 1022ms | ✓ 1573ms | ✓ 835ms | http |
| 38.145.220.175:8444 | ✓ 612ms | ✓ 1131ms | ✓ 1651ms | ✓ 1083ms | ✓ 866ms | http |
| 38.145.220.36:8444 | ✓ 615ms | 否 | ✓ 713ms | ✓ 1572ms | ✓ 802ms | http |
| 45.136.130.197:8452 | ✓ 610ms | ✓ 1547ms | ✓ 680ms | ✓ 1621ms | ✓ 799ms | http |
| 38.145.220.39:8450 | ✓ 763ms | 否 | ✓ 933ms | ✓ 1488ms | 否 | http |
| 45.136.130.198:8452 | ✓ 610ms | ✓ 1981ms | ✓ 771ms | ✓ 1600ms | ✓ 1175ms | http |
| 45.136.131.25:8445 | ✓ 611ms | ✓ 1620ms | ✓ 765ms | 否 | ✓ 772ms | http |
| 38.145.208.224:8445 | ✓ 1172ms | ✓ 920ms | ✓ 622ms | ✓ 1728ms | ✓ 1149ms | http |
| 38.145.208.219:8445 | ✓ 1173ms | ✓ 1058ms | ✓ 658ms | ✓ 1598ms | ✓ 1134ms | http |
| 38.145.208.240:8448 | ✓ 1866ms | ✓ 1152ms | ✓ 619ms | 否 | ✓ 718ms | http |
| 38.145.208.247:8448 | ✓ 1883ms | ✓ 1139ms | ✓ 625ms | 否 | ✓ 746ms | http |
| 38.145.208.243:8448 | ✓ 1869ms | ✓ 1195ms | ✓ 732ms | 否 | ✓ 782ms | http |
| 180.127.149.225:1080 | ✓ 1033ms | 否 | ✓ 1694ms | 否 | ✓ 1085ms | http |
| 38.145.208.226:8453 | ✓ 1792ms | 否 | ✓ 807ms | ✓ 1945ms | ✓ 718ms | http |
| 39.98.86.246:8118 | 否 | 否 | ✓ 1419ms | ✓ 1298ms | ✓ 1042ms | http |
| 38.145.208.169:8444 | ✓ 1814ms | ✓ 931ms | ✓ 547ms | ✓ 1287ms | ✓ 1504ms | http |
| 123.57.0.163:8888 | ✓ 1591ms | ✓ 1850ms | 否 | 否 | ✓ 1573ms | http |
| 121.230.8.213:1080 | 否 | ✓ 1578ms | ✓ 1170ms | 否 | ✓ 1437ms | http |
| 52.163.56.148:80 | ✓ 1409ms | ✓ 1498ms | ✓ 982ms | ✓ 1299ms | 否 | http |
| 197.164.101.14:1976 | ✓ 1463ms | 否 | ✓ 1286ms | 否 | ✓ 1690ms | http |
| 45.136.131.65:8451 | ✓ 1744ms | 否 | ✓ 851ms | 否 | ✓ 1297ms | http |
| 106.117.208.101:7890 | ✓ 1097ms | 否 | ✓ 1459ms | ✓ 1406ms | 否 | http |
| 47.79.40.38:55000 | ✓ 1213ms | 否 | ✓ 1272ms | ✓ 1184ms | 否 | http |

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
