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

最后更新：2026-04-07 11:58:30 UTC（2026-04-07 19:58:30 UTC+8）

**代理总数：145**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 145 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 145 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 670ms | ✓ 1036ms | ✓ 660ms | ✓ 868ms | ✓ 688ms | http |
| 113.160.132.26:8080 | ✓ 1410ms | ✓ 1266ms | ✓ 985ms | ✓ 1178ms | ✓ 923ms | http |
| 159.223.71.162:8080 | ✓ 1425ms | 否 | ✓ 858ms | ✓ 1176ms | ✓ 868ms | http |
| 159.223.71.162:443 | ✓ 1428ms | 否 | ✓ 874ms | ✓ 1193ms | ✓ 873ms | http |
| 147.161.210.140:8800 | ✓ 1254ms | 否 | ✓ 1207ms | ✓ 1266ms | ✓ 1305ms | http |
| 167.103.115.102:8800 | ✓ 1101ms | 否 | ✓ 1182ms | ✓ 1288ms | ✓ 1359ms | http |
| 159.223.213.91:8888 | ✓ 771ms | 否 | 否 | ✓ 1425ms | ✓ 1289ms | http |
| 111.227.254.12:22222 | ✓ 1083ms | ✓ 1374ms | ✓ 1113ms | ✓ 1418ms | ✓ 1072ms | http |
| 167.103.34.108:8800 | ✓ 1462ms | 否 | ✓ 1217ms | 否 | ✓ 1575ms | http |
| 111.227.254.11:22222 | ✓ 1074ms | 否 | ✓ 1643ms | ✓ 1305ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1288ms | ✓ 1276ms | ✓ 1803ms | 否 | ✓ 1411ms | http |
| 45.88.0.98:3128 | ✓ 1687ms | 否 | ✓ 1260ms | ✓ 1531ms | ✓ 1161ms | http |
| 45.88.0.117:3128 | 否 | ✓ 1481ms | ✓ 1464ms | 否 | ✓ 1149ms | http |
| 35.225.22.61:80 | ✓ 485ms | 否 | ✓ 436ms | 否 | ✓ 923ms | http |
| 38.34.179.39:8452 | ✓ 456ms | ✓ 822ms | ✓ 619ms | ✓ 1688ms | ✓ 589ms | http |
| 45.136.130.169:8444 | ✓ 1293ms | ✓ 1711ms | ✓ 219ms | ✓ 823ms | ✓ 600ms | http |
| 45.136.130.167:8448 | ✓ 682ms | ✓ 1311ms | ✓ 635ms | ✓ 923ms | ✓ 691ms | http |
| 38.34.179.74:8447 | ✓ 476ms | ✓ 1081ms | ✓ 881ms | ✓ 1040ms | ✓ 624ms | http |
| 38.34.179.75:8447 | ✓ 501ms | ✓ 1126ms | ✓ 853ms | ✓ 1079ms | ✓ 645ms | http |
| 45.136.131.25:8453 | ✓ 1148ms | ✓ 1421ms | ✓ 787ms | ✓ 1443ms | ✓ 1106ms | http |
| 45.136.131.35:8452 | ✓ 1135ms | 否 | ✓ 387ms | ✓ 1276ms | ✓ 1108ms | http |
| 38.34.183.47:8452 | ✓ 423ms | ✓ 1795ms | ✓ 1285ms | ✓ 768ms | ✓ 591ms | http |
| 38.34.179.76:8444 | ✓ 499ms | ✓ 1096ms | ✓ 1040ms | ✓ 1095ms | ✓ 835ms | http |
| 38.34.179.79:8449 | ✓ 589ms | ✓ 1362ms | ✓ 983ms | ✓ 1173ms | ✓ 1032ms | http |
| 34.96.238.40:8080 | ✓ 1449ms | ✓ 1550ms | 否 | ✓ 1092ms | 否 | http |
| 38.34.183.225:8450 | ✓ 752ms | 否 | ✓ 847ms | ✓ 1617ms | ✓ 1399ms | http |
| 38.145.220.49:8444 | ✓ 439ms | ✓ 1188ms | 否 | ✓ 1386ms | ✓ 1472ms | http |
| 45.136.131.68:8446 | ✓ 434ms | ✓ 1198ms | ✓ 1013ms | ✓ 1877ms | ✓ 1067ms | http |
| 38.145.208.209:8447 | ✓ 1228ms | 否 | ✓ 1199ms | ✓ 1610ms | 否 | http |
| 130.61.30.221:8080 | ✓ 1030ms | 否 | ✓ 1457ms | 否 | ✓ 1819ms | http |
| 111.227.254.9:22222 | ✓ 1066ms | ✓ 1335ms | ✓ 1018ms | 否 | ✓ 1115ms | http |
| 38.34.179.72:8452 | ✓ 1010ms | ✓ 1767ms | ✓ 1449ms | ✓ 1318ms | ✓ 748ms | http |
| 5.104.87.17:8051 | ✓ 1154ms | 否 | ✓ 1146ms | ✓ 1673ms | ✓ 1647ms | http |
| 38.34.179.56:8450 | ✓ 1592ms | 否 | ✓ 1671ms | ✓ 1095ms | ✓ 579ms | http |
| 167.103.144.127:8800 | ✓ 1895ms | 否 | ✓ 1376ms | ✓ 1806ms | ✓ 1654ms | http |
| 38.34.179.18:8451 | ✓ 1119ms | 否 | ✓ 897ms | 否 | ✓ 1307ms | http |
| 111.227.254.10:22222 | ✓ 976ms | ✓ 1275ms | ✓ 1555ms | 否 | ✓ 1053ms | http |
| 45.136.131.66:8448 | ✓ 424ms | ✓ 1549ms | ✓ 1399ms | ✓ 1458ms | ✓ 1275ms | http |
| 45.136.131.67:8448 | ✓ 471ms | 否 | ✓ 1482ms | ✓ 1446ms | ✓ 1296ms | http |
| 212.58.132.5:8888 | ✓ 1865ms | 否 | ✓ 1894ms | ✓ 1728ms | ✓ 1874ms | http |
| 45.167.124.52:8080 | ✓ 1184ms | 否 | ✓ 641ms | ✓ 1656ms | ✓ 1427ms | http |
| 167.103.31.122:8800 | ✓ 1393ms | 否 | ✓ 1377ms | ✓ 1702ms | 否 | http |
| 45.88.0.99:3128 | ✓ 553ms | 否 | ✓ 626ms | ✓ 1849ms | 否 | http |
| 45.88.0.116:3128 | ✓ 600ms | ✓ 1360ms | ✓ 548ms | 否 | 否 | http |
| 45.88.0.115:3128 | ✓ 551ms | ✓ 1467ms | ✓ 519ms | 否 | 否 | http |
| 152.32.132.190:7890 | ✓ 1762ms | 否 | 否 | ✓ 1965ms | ✓ 1335ms | http |
| 38.34.179.77:8446 | ✓ 769ms | ✓ 1772ms | ✓ 254ms | ✓ 784ms | ✓ 930ms | http |
| 38.34.179.88:8446 | ✓ 1174ms | ✓ 693ms | ✓ 551ms | ✓ 1294ms | ✓ 1594ms | http |
| 147.161.239.240:8800 | ✓ 1028ms | ✓ 1862ms | ✓ 1477ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 1631ms | ✓ 1158ms | 否 | ✓ 1269ms | ✓ 894ms | http |
| 109.107.179.140:8090 | ✓ 1038ms | 否 | ✓ 1591ms | 否 | ✓ 1719ms | http |
| 45.167.125.21:999 | ✓ 803ms | 否 | ✓ 1545ms | ✓ 1819ms | ✓ 1499ms | http |
| 165.22.57.158:8080 | ✓ 1249ms | 否 | ✓ 1680ms | ✓ 1923ms | ✓ 844ms | http |
| 178.128.24.162:8080 | ✓ 785ms | 否 | 否 | ✓ 1128ms | ✓ 1020ms | http |
| 103.252.89.130:8080 | ✓ 1061ms | 否 | ✓ 861ms | ✓ 1844ms | 否 | http |
| 159.223.225.118:8888 | ✓ 1768ms | 否 | ✓ 659ms | ✓ 1868ms | ✓ 1759ms | http |
| 38.145.208.175:8447 | 否 | ✓ 1465ms | ✓ 515ms | ✓ 862ms | ✓ 1541ms | http |
| 38.145.208.211:8453 | ✓ 1503ms | ✓ 1169ms | ✓ 189ms | ✓ 1013ms | ✓ 694ms | http |
| 150.249.255.91:3128 | ✓ 643ms | 否 | ✓ 598ms | ✓ 933ms | 否 | http |
| 45.136.130.176:8451 | ✓ 728ms | ✓ 1969ms | ✓ 340ms | ✓ 760ms | ✓ 651ms | http |
| 38.34.179.86:8452 | ✓ 275ms | ✓ 1580ms | ✓ 462ms | ✓ 1208ms | ✓ 581ms | http |
| 150.241.71.15:1080 | ✓ 1086ms | ✓ 1964ms | ✓ 1115ms | ✓ 1740ms | ✓ 1306ms | http |
| 150.107.140.238:3128 | ✓ 1021ms | 否 | 否 | ✓ 1217ms | ✓ 1150ms | http |
| 59.46.216.131:30001 | ✓ 994ms | 否 | ✓ 1891ms | ✓ 1370ms | ✓ 1114ms | http |
| 38.180.2.107:3128 | ✓ 1040ms | ✓ 1777ms | ✓ 1662ms | 否 | ✓ 1757ms | http |
| 62.113.119.14:8080 | ✓ 956ms | ✓ 1615ms | ✓ 1330ms | ✓ 1632ms | ✓ 1171ms | http |
| 43.153.28.68:3128 | ✓ 442ms | 否 | ✓ 538ms | ✓ 1079ms | 否 | http |
| 38.34.179.51:8449 | ✓ 1611ms | ✓ 1402ms | ✓ 1621ms | 否 | ✓ 840ms | http |
| 45.88.0.114:3128 | 否 | 否 | ✓ 1624ms | ✓ 1488ms | ✓ 1287ms | http |
| 38.34.183.224:8448 | 否 | ✓ 1168ms | ✓ 163ms | ✓ 772ms | ✓ 598ms | http |
| 38.34.179.24:8447 | 否 | ✓ 700ms | ✓ 249ms | ✓ 1074ms | ✓ 1036ms | http |
| 38.34.179.135:8450 | 否 | ✓ 993ms | ✓ 708ms | ✓ 1486ms | ✓ 560ms | http |
| 38.34.179.178:8444 | 否 | 否 | ✓ 156ms | ✓ 736ms | ✓ 623ms | http |
| 38.34.179.175:8445 | 否 | ✓ 1522ms | ✓ 266ms | ✓ 804ms | ✓ 689ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1075ms | ✓ 1964ms | ✓ 1505ms | ✓ 997ms | http |
| 38.34.179.204:8447 | 否 | ✓ 1452ms | ✓ 1193ms | ✓ 754ms | ✓ 832ms | http |
| 45.136.131.48:8446 | 否 | ✓ 1905ms | ✓ 1660ms | ✓ 1177ms | 否 | http |
| 103.39.51.207:8080 | ✓ 1331ms | 否 | 否 | ✓ 1435ms | ✓ 1366ms | http |
| 103.143.81.156:7890 | ✓ 744ms | 否 | ✓ 1068ms | ✓ 937ms | ✓ 755ms | http |
| 161.35.70.36:8888 | ✓ 641ms | 否 | ✓ 1274ms | ✓ 1972ms | ✓ 1517ms | http |
| 180.103.19.151:1080 | 否 | ✓ 1654ms | ✓ 1667ms | 否 | ✓ 1664ms | http |
| 103.135.102.161:8081 | ✓ 1837ms | 否 | ✓ 1062ms | ✓ 1450ms | 否 | http |
| 45.136.130.171:8450 | 否 | ✓ 1975ms | ✓ 272ms | ✓ 1866ms | ✓ 574ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1041ms | ✓ 1333ms | ✓ 1896ms | http |
| 38.34.179.155:8453 | ✓ 633ms | ✓ 918ms | ✓ 239ms | ✓ 767ms | ✓ 587ms | http |
| 38.34.179.57:8448 | ✓ 936ms | ✓ 1644ms | ✓ 163ms | ✓ 769ms | ✓ 603ms | http |
| 38.34.179.80:8452 | ✓ 568ms | ✓ 732ms | ✓ 488ms | ✓ 1118ms | ✓ 1039ms | http |
| 38.34.179.21:8443 | ✓ 551ms | ✓ 1924ms | ✓ 825ms | ✓ 735ms | ✓ 707ms | http |
| 38.34.179.97:8446 | ✓ 773ms | 否 | ✓ 161ms | ✓ 742ms | ✓ 639ms | http |
| 45.136.130.181:8445 | ✓ 549ms | ✓ 1482ms | ✓ 1661ms | ✓ 1417ms | ✓ 582ms | http |
| 45.136.130.182:8450 | ✓ 550ms | ✓ 1728ms | ✓ 1410ms | ✓ 811ms | ✓ 967ms | http |
| 38.145.220.22:8447 | ✓ 1306ms | 否 | ✓ 1247ms | ✓ 817ms | 否 | http |
| 38.34.179.50:8444 | ✓ 231ms | ✓ 727ms | ✓ 360ms | ✓ 1185ms | ✓ 588ms | http |
| 38.34.183.11:8444 | ✓ 589ms | ✓ 730ms | ✓ 705ms | ✓ 871ms | ✓ 591ms | http |
| 38.34.183.8:8444 | ✓ 597ms | ✓ 765ms | ✓ 662ms | ✓ 908ms | ✓ 581ms | http |
| 38.145.218.235:8444 | ✓ 605ms | ✓ 1593ms | ✓ 304ms | ✓ 864ms | ✓ 579ms | http |
| 45.136.131.57:8450 | ✓ 194ms | ✓ 1034ms | ✓ 311ms | 否 | ✓ 926ms | http |
| 38.34.179.202:8448 | ✓ 615ms | ✓ 1683ms | ✓ 159ms | ✓ 787ms | ✓ 577ms | http |
| 38.145.208.182:8452 | ✓ 926ms | 否 | ✓ 144ms | ✓ 761ms | ✓ 863ms | http |
| 45.136.131.62:8451 | ✓ 1438ms | ✓ 1112ms | ✓ 372ms | ✓ 900ms | ✓ 911ms | http |

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
