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

最后更新：2026-03-29 18:26:05 UTC（2026-03-30 02:26:05 UTC+8）

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
| 208.87.243.199:7878 | ✓ 500ms | ✓ 807ms | ✓ 1071ms | ✓ 870ms | ✓ 616ms | http |
| 39.185.46.193:5911 | ✓ 709ms | ✓ 863ms | ✓ 726ms | ✓ 937ms | ✓ 806ms | http |
| 43.99.54.236:5555 | ✓ 715ms | ✓ 1198ms | ✓ 672ms | ✓ 837ms | ✓ 673ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 834ms | ✓ 907ms | ✓ 1107ms | http |
| 147.161.239.240:8800 | ✓ 1220ms | ✓ 1923ms | ✓ 1164ms | ✓ 1535ms | ✓ 1360ms | http |
| 167.103.115.102:8800 | ✓ 938ms | ✓ 1662ms | ✓ 1053ms | ✓ 1248ms | ✓ 1802ms | http |
| 103.84.95.54:7890 | ✓ 755ms | 否 | ✓ 748ms | ✓ 1598ms | 否 | http |
| 95.213.217.168:52004 | ✓ 1269ms | ✓ 1729ms | ✓ 1585ms | ✓ 1869ms | ✓ 1430ms | http |
| 42.96.16.158:1311 | ✓ 1908ms | 否 | ✓ 1566ms | ✓ 1277ms | ✓ 1412ms | http |
| 167.103.34.108:8800 | ✓ 1538ms | 否 | ✓ 1491ms | 否 | ✓ 1589ms | http |
| 113.160.132.26:8080 | ✓ 1590ms | ✓ 1659ms | ✓ 1159ms | 否 | ✓ 1025ms | http |
| 180.250.219.58:53281 | ✓ 1555ms | 否 | ✓ 1514ms | 否 | ✓ 1900ms | http |
| 35.225.22.61:80 | ✓ 395ms | 否 | ✓ 1301ms | ✓ 1197ms | ✓ 1024ms | http |
| 167.103.144.127:8800 | ✓ 961ms | ✓ 1829ms | ✓ 1409ms | ✓ 1487ms | ✓ 1610ms | http |
| 64.227.76.27:1080 | ✓ 712ms | ✓ 1678ms | ✓ 1667ms | ✓ 1921ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1664ms | 否 | ✓ 1758ms | 否 | ✓ 1958ms | http |
| 103.113.70.189:1081 | ✓ 1033ms | ✓ 1089ms | 否 | ✓ 1137ms | 否 | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 657ms | ✓ 1043ms | ✓ 1156ms | http |
| 198.59.68.130:3128 | ✓ 988ms | ✓ 1277ms | ✓ 1691ms | ✓ 1419ms | ✓ 1315ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1185ms | ✓ 1971ms | ✓ 1926ms | http |
| 120.92.212.16:8890 | ✓ 1991ms | ✓ 1248ms | ✓ 1777ms | 否 | ✓ 1204ms | http |
| 59.46.216.131:30001 | ✓ 939ms | ✓ 1352ms | ✓ 977ms | ✓ 1341ms | 否 | http |
| 38.34.179.61:8444 | ✓ 1734ms | ✓ 881ms | ✓ 1998ms | ✓ 1575ms | ✓ 636ms | http |
| 101.43.127.100:8877 | ✓ 897ms | ✓ 1151ms | ✓ 919ms | ✓ 1184ms | ✓ 1862ms | http |
| 1.231.81.166:3128 | ✓ 1501ms | ✓ 1194ms | 否 | ✓ 1470ms | ✓ 857ms | http |
| 5.102.109.41:999 | ✓ 1113ms | ✓ 1441ms | 否 | ✓ 1304ms | ✓ 1478ms | http |
| 8.219.97.248:80 | ✓ 1480ms | 否 | ✓ 1173ms | ✓ 1681ms | ✓ 1231ms | http |
| 38.34.179.164:8448 | ✓ 355ms | ✓ 939ms | ✓ 1285ms | ✓ 1687ms | ✓ 1260ms | http |
| 45.12.151.226:2829 | ✓ 1273ms | ✓ 1707ms | 否 | 否 | ✓ 1388ms | http |
| 101.47.73.135:3128 | ✓ 1741ms | 否 | 否 | ✓ 1379ms | ✓ 1146ms | http |
| 116.80.49.161:3172 | 否 | 否 | ✓ 1522ms | ✓ 1883ms | ✓ 1679ms | http |
| 38.34.183.225:8444 | ✓ 214ms | ✓ 732ms | ✓ 414ms | ✓ 858ms | ✓ 742ms | http |
| 193.233.22.29:10808 | ✓ 753ms | 否 | ✓ 628ms | ✓ 1458ms | 否 | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1679ms | ✓ 1520ms | ✓ 1894ms | http |
| 177.234.217.88:999 | ✓ 1341ms | 否 | ✓ 1055ms | ✓ 1845ms | ✓ 1566ms | http |
| 91.238.123.111:8000 | ✓ 515ms | ✓ 1495ms | ✓ 1378ms | ✓ 1673ms | ✓ 1209ms | http |
| 91.238.123.230:8000 | ✓ 535ms | ✓ 1450ms | ✓ 1382ms | ✓ 1678ms | ✓ 1198ms | http |
| 38.34.179.21:8446 | 否 | ✓ 1443ms | ✓ 1515ms | ✓ 781ms | ✓ 819ms | http |
| 46.39.105.157:8080 | ✓ 1662ms | 否 | ✓ 1168ms | 否 | ✓ 1771ms | http |
| 103.183.10.169:3125 | ✓ 1979ms | 否 | ✓ 1701ms | ✓ 1707ms | ✓ 1430ms | http |
| 209.126.84.232:8888 | 否 | 否 | ✓ 1315ms | ✓ 1593ms | ✓ 1323ms | http |
| 45.144.28.81:10808 | ✓ 589ms | ✓ 1251ms | ✓ 1371ms | ✓ 1466ms | ✓ 1312ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1827ms | ✓ 1368ms | ✓ 1397ms | http |
| 38.145.220.60:8444 | ✓ 1354ms | ✓ 710ms | ✓ 721ms | ✓ 1117ms | ✓ 1481ms | http |
| 38.145.208.204:8445 | ✓ 1126ms | 否 | ✓ 493ms | ✓ 1057ms | ✓ 1905ms | http |
| 38.145.208.224:8446 | ✓ 941ms | 否 | ✓ 680ms | ✓ 1007ms | 否 | http |
| 38.145.208.169:8446 | ✓ 345ms | ✓ 1162ms | 否 | ✓ 811ms | ✓ 727ms | http |
| 38.145.208.171:8453 | ✓ 366ms | ✓ 1345ms | ✓ 1891ms | ✓ 822ms | ✓ 1301ms | http |
| 38.34.179.192:8451 | ✓ 391ms | ✓ 1276ms | 否 | ✓ 850ms | ✓ 899ms | http |
| 38.145.208.195:8453 | ✓ 345ms | ✓ 798ms | ✓ 1428ms | ✓ 1762ms | ✓ 983ms | http |
| 38.145.208.193:8451 | ✓ 346ms | ✓ 810ms | ✓ 977ms | 否 | ✓ 891ms | http |
| 180.103.19.226:1080 | ✓ 1064ms | 否 | ✓ 1251ms | ✓ 1355ms | ✓ 1083ms | http |
| 195.123.209.48:3128 | ✓ 1228ms | ✓ 1977ms | ✓ 1244ms | ✓ 1965ms | ✓ 1783ms | http |
| 38.145.208.191:8446 | ✓ 348ms | ✓ 875ms | ✓ 1320ms | ✓ 1772ms | ✓ 1361ms | http |
| 121.230.8.34:1080 | ✓ 1597ms | ✓ 1845ms | ✓ 1289ms | ✓ 1818ms | ✓ 1120ms | http |
| 38.34.179.100:8452 | 否 | ✓ 1032ms | ✓ 1122ms | 否 | ✓ 810ms | http |
| 38.145.208.204:8451 | ✓ 1268ms | 否 | ✓ 579ms | ✓ 1904ms | ✓ 1542ms | http |
| 38.145.220.198:8451 | ✓ 1361ms | ✓ 1179ms | ✓ 1683ms | ✓ 1906ms | 否 | http |
| 45.140.147.82:1081 | 否 | ✓ 1241ms | 否 | ✓ 1335ms | ✓ 1972ms | http |
| 121.230.8.213:1080 | 否 | ✓ 1205ms | ✓ 1336ms | ✓ 1561ms | ✓ 1116ms | http |
| 82.146.58.184:1080 | ✓ 1551ms | 否 | ✓ 756ms | 否 | ✓ 1877ms | http |
| 116.80.95.238:7777 | ✓ 1516ms | 否 | ✓ 1509ms | 否 | ✓ 1660ms | http |
| 45.136.198.40:3128 | ✓ 821ms | ✓ 1682ms | ✓ 1912ms | 否 | ✓ 1879ms | http |
| 38.145.208.172:8448 | ✓ 1274ms | 否 | ✓ 1647ms | ✓ 965ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1935ms | 否 | 否 | ✓ 1485ms | ✓ 1331ms | http |
| 116.80.65.79:3172 | ✓ 1495ms | 否 | 否 | ✓ 1848ms | ✓ 1684ms | http |
| 62.113.119.14:8080 | ✓ 741ms | 否 | ✓ 844ms | ✓ 1772ms | ✓ 1365ms | http |
| 218.60.0.214:80 | ✓ 1027ms | ✓ 1358ms | ✓ 1064ms | ✓ 1350ms | ✓ 1021ms | http |
| 194.163.183.242:3128 | ✓ 1135ms | ✓ 1936ms | ✓ 1635ms | ✓ 1722ms | ✓ 1284ms | http |
| 114.237.77.220:1080 | ✓ 1013ms | ✓ 1962ms | 否 | 否 | ✓ 1545ms | http |
| 106.75.15.167:7890 | ✓ 1476ms | ✓ 1188ms | ✓ 932ms | 否 | ✓ 974ms | http |
| 180.103.19.233:1080 | 否 | ✓ 1416ms | ✓ 1205ms | ✓ 1627ms | ✓ 1126ms | http |
| 107.178.115.140:3128 | ✓ 503ms | ✓ 793ms | ✓ 1205ms | ✓ 1047ms | ✓ 857ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1396ms | ✓ 1091ms | ✓ 1546ms | ✓ 942ms | http |
| 37.59.110.73:80 | ✓ 1236ms | ✓ 1875ms | ✓ 1563ms | 否 | 否 | http |
| 20.120.225.109:3128 | ✓ 419ms | ✓ 1194ms | ✓ 699ms | ✓ 1136ms | ✓ 853ms | http |
| 202.129.206.239:3128 | ✓ 1435ms | 否 | ✓ 1583ms | ✓ 1606ms | ✓ 1506ms | http |

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
