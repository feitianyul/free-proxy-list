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

最后更新：2026-04-01 23:41:46 UTC（2026-04-02 07:41:46 UTC+8）

**代理总数：132**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 132 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 132 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 763ms | ✓ 938ms | ✓ 634ms | ✓ 801ms | ✓ 634ms | http |
| 147.161.210.140:8800 | ✓ 757ms | 否 | ✓ 814ms | ✓ 1285ms | ✓ 877ms | http |
| 1.231.81.166:3128 | ✓ 996ms | ✓ 915ms | ✓ 1842ms | ✓ 1143ms | ✓ 833ms | http |
| 133.242.138.34:8100 | ✓ 1048ms | ✓ 1720ms | ✓ 866ms | ✓ 1220ms | ✓ 1005ms | http |
| 159.223.71.162:443 | ✓ 817ms | 否 | ✓ 821ms | ✓ 1029ms | ✓ 821ms | http |
| 38.34.179.40:8446 | ✓ 806ms | ✓ 1329ms | ✓ 1483ms | ✓ 832ms | ✓ 835ms | http |
| 167.103.115.102:8800 | ✓ 920ms | ✓ 1690ms | ✓ 1251ms | ✓ 1534ms | ✓ 1126ms | http |
| 113.160.132.26:8080 | ✓ 1430ms | ✓ 1303ms | ✓ 1410ms | ✓ 1311ms | ✓ 985ms | http |
| 42.96.16.158:1311 | ✓ 1420ms | 否 | ✓ 1128ms | ✓ 1577ms | ✓ 994ms | http |
| 83.219.250.8:62920 | ✓ 885ms | 否 | ✓ 1714ms | 否 | ✓ 1482ms | http |
| 212.58.132.5:8888 | ✓ 1074ms | ✓ 1984ms | ✓ 1434ms | ✓ 1826ms | ✓ 1222ms | http |
| 95.213.217.168:52004 | ✓ 1101ms | ✓ 1806ms | ✓ 1593ms | 否 | ✓ 1792ms | http |
| 167.103.34.108:8800 | ✓ 1413ms | 否 | ✓ 1358ms | 否 | ✓ 1394ms | http |
| 45.167.124.52:8080 | ✓ 1180ms | 否 | ✓ 1549ms | 否 | ✓ 1749ms | http |
| 208.87.243.199:7878 | ✓ 255ms | ✓ 649ms | ✓ 229ms | ✓ 666ms | ✓ 491ms | http |
| 38.34.179.86:8452 | ✓ 234ms | ✓ 867ms | ✓ 689ms | ✓ 1009ms | ✓ 504ms | http |
| 38.34.179.150:8449 | ✓ 1027ms | ✓ 668ms | ✓ 686ms | ✓ 943ms | ✓ 1178ms | http |
| 38.34.179.51:8449 | ✓ 1103ms | ✓ 1080ms | ✓ 125ms | ✓ 882ms | ✓ 851ms | http |
| 203.80.138.81:50000 | ✓ 859ms | ✓ 1036ms | ✓ 817ms | ✓ 962ms | ✓ 814ms | http |
| 38.34.179.54:8453 | ✓ 1381ms | ✓ 693ms | ✓ 116ms | ✓ 928ms | ✓ 789ms | http |
| 38.34.179.88:8446 | ✓ 1071ms | ✓ 1695ms | ✓ 85ms | ✓ 704ms | ✓ 778ms | http |
| 38.34.183.130:8452 | ✓ 1122ms | ✓ 1747ms | ✓ 948ms | ✓ 747ms | ✓ 1010ms | http |
| 34.96.238.40:8080 | ✓ 800ms | ✓ 1712ms | ✓ 1464ms | 否 | ✓ 1671ms | http |
| 167.103.144.127:8800 | ✓ 1388ms | ✓ 1692ms | ✓ 953ms | ✓ 1734ms | ✓ 1417ms | http |
| 180.250.219.58:53281 | ✓ 1752ms | 否 | ✓ 1430ms | ✓ 1902ms | ✓ 1815ms | http |
| 35.225.22.61:80 | ✓ 485ms | ✓ 1301ms | ✓ 329ms | 否 | 否 | http |
| 45.12.151.226:2829 | ✓ 978ms | ✓ 1855ms | ✓ 1180ms | ✓ 1849ms | ✓ 1397ms | http |
| 167.103.31.122:8800 | ✓ 1283ms | 否 | ✓ 1286ms | 否 | ✓ 1474ms | http |
| 149.62.191.202:3128 | ✓ 1003ms | ✓ 1979ms | ✓ 1031ms | 否 | ✓ 1378ms | http |
| 190.12.150.244:999 | ✓ 1065ms | ✓ 1886ms | ✓ 1049ms | 否 | ✓ 1537ms | http |
| 223.16.170.103:80 | 否 | ✓ 1633ms | ✓ 1172ms | 否 | ✓ 1054ms | http |
| 38.145.208.242:8451 | ✓ 712ms | ✓ 833ms | ✓ 348ms | ✓ 1323ms | 否 | http |
| 38.34.183.224:8448 | ✓ 890ms | ✓ 775ms | ✓ 1073ms | 否 | 否 | http |
| 38.145.220.102:8453 | ✓ 951ms | ✓ 645ms | ✓ 144ms | ✓ 838ms | 否 | http |
| 167.160.191.204:6005 | ✓ 680ms | ✓ 1462ms | ✓ 1630ms | 否 | ✓ 1132ms | http |
| 38.34.179.39:8452 | ✓ 191ms | ✓ 769ms | ✓ 1011ms | 否 | 否 | http |
| 38.34.179.98:8451 | ✓ 688ms | ✓ 977ms | ✓ 82ms | ✓ 656ms | 否 | http |
| 38.34.179.47:8452 | ✓ 834ms | ✓ 688ms | ✓ 202ms | ✓ 935ms | ✓ 1275ms | http |
| 45.136.130.188:8449 | ✓ 1282ms | ✓ 1687ms | ✓ 889ms | 否 | ✓ 1192ms | http |
| 38.34.179.186:8444 | ✓ 621ms | 否 | ✓ 97ms | ✓ 710ms | ✓ 846ms | http |
| 120.92.212.16:7890 | ✓ 1485ms | ✓ 1159ms | 否 | ✓ 960ms | ✓ 958ms | http |
| 38.145.208.172:8448 | ✓ 730ms | ✓ 622ms | ✓ 954ms | ✓ 1401ms | ✓ 526ms | http |
| 101.43.127.100:8877 | ✓ 647ms | ✓ 831ms | ✓ 700ms | ✓ 849ms | ✓ 716ms | http |
| 45.136.130.170:8448 | ✓ 699ms | ✓ 1628ms | ✓ 1066ms | ✓ 755ms | ✓ 667ms | http |
| 45.136.130.173:8448 | ✓ 698ms | ✓ 1735ms | ✓ 961ms | ✓ 772ms | ✓ 659ms | http |
| 167.160.184.231:6005 | ✓ 1545ms | ✓ 1450ms | ✓ 1123ms | ✓ 1603ms | ✓ 1030ms | http |
| 160.250.5.22:1 | ✓ 1358ms | 否 | ✓ 1420ms | ✓ 1301ms | ✓ 959ms | http |
| 147.161.239.240:8800 | ✓ 1252ms | ✓ 1767ms | ✓ 1092ms | ✓ 1870ms | ✓ 1686ms | http |
| 120.92.212.16:8890 | 否 | ✓ 936ms | 否 | ✓ 1784ms | ✓ 1552ms | http |
| 59.46.216.131:30001 | ✓ 1760ms | ✓ 1061ms | ✓ 1904ms | ✓ 1069ms | ✓ 855ms | http |
| 177.234.217.88:999 | ✓ 1590ms | 否 | ✓ 1811ms | 否 | ✓ 1902ms | http |
| 217.182.195.221:30003 | ✓ 959ms | 否 | ✓ 1931ms | 否 | ✓ 1553ms | http |
| 38.34.179.35:8448 | ✓ 454ms | ✓ 737ms | ✓ 82ms | ✓ 692ms | ✓ 553ms | http |
| 38.34.179.27:8451 | 否 | ✓ 1245ms | ✓ 856ms | 否 | ✓ 1086ms | http |
| 165.232.146.249:3128 | ✓ 1084ms | ✓ 1124ms | 否 | 否 | ✓ 1570ms | http |
| 148.153.56.51:80 | ✓ 296ms | ✓ 632ms | ✓ 689ms | ✓ 878ms | ✓ 616ms | http |
| 138.197.68.35:4857 | 否 | ✓ 1499ms | ✓ 559ms | ✓ 1455ms | 否 | http |
| 1.225.116.115:1080 | 否 | ✓ 1702ms | ✓ 1078ms | ✓ 1737ms | ✓ 1067ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1470ms | ✓ 1488ms | ✓ 1627ms | http |
| 47.101.159.19:8899 | ✓ 1853ms | ✓ 1089ms | ✓ 1570ms | ✓ 1153ms | 否 | http |
| 49.7.179.70:3333 | ✓ 1107ms | ✓ 1636ms | ✓ 1755ms | ✓ 1411ms | ✓ 847ms | http |
| 38.34.179.14:8450 | ✓ 337ms | ✓ 1220ms | ✓ 1571ms | ✓ 1406ms | ✓ 1720ms | http |
| 72.11.150.178:6005 | ✓ 993ms | ✓ 1617ms | ✓ 1270ms | ✓ 1437ms | 否 | http |
| 45.136.130.169:8444 | ✓ 1380ms | 否 | ✓ 761ms | ✓ 830ms | ✓ 1147ms | http |
| 38.145.218.161:8443 | ✓ 628ms | ✓ 938ms | ✓ 90ms | ✓ 686ms | ✓ 574ms | http |
| 38.34.183.47:8452 | ✓ 520ms | ✓ 1169ms | ✓ 159ms | ✓ 720ms | ✓ 590ms | http |
| 34.101.184.164:3128 | ✓ 1412ms | 否 | ✓ 785ms | ✓ 1704ms | ✓ 1111ms | http |
| 157.230.220.25:4857 | ✓ 966ms | 否 | 否 | ✓ 1959ms | ✓ 1320ms | http |
| 23.224.193.43:3128 | ✓ 860ms | 否 | 否 | ✓ 1724ms | ✓ 1734ms | http |
| 85.208.108.43:10808 | ✓ 992ms | 否 | ✓ 1602ms | ✓ 1705ms | ✓ 1347ms | http |
| 20.120.225.109:3128 | ✓ 641ms | ✓ 1220ms | ✓ 942ms | ✓ 1170ms | ✓ 766ms | http |
| 180.103.19.233:1080 | ✓ 867ms | ✓ 957ms | ✓ 963ms | ✓ 1011ms | ✓ 1018ms | http |
| 23.224.193.44:3128 | ✓ 1128ms | 否 | ✓ 1033ms | ✓ 1330ms | ✓ 1066ms | http |
| 157.230.38.173:3128 | ✓ 1170ms | 否 | ✓ 744ms | ✓ 1015ms | ✓ 802ms | http |
| 23.224.193.45:3128 | ✓ 1273ms | ✓ 1235ms | ✓ 1832ms | ✓ 1310ms | ✓ 1274ms | http |
| 150.31.45.65:8118 | ✓ 769ms | 否 | ✓ 1671ms | 否 | ✓ 793ms | http |
| 120.92.211.211:7890 | ✓ 1715ms | 否 | ✓ 653ms | 否 | ✓ 888ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1290ms | ✓ 720ms | ✓ 1936ms | ✓ 734ms | http |
| 62.113.119.14:8080 | ✓ 1529ms | ✓ 1908ms | ✓ 1455ms | ✓ 1767ms | 否 | http |
| 117.86.6.27:1080 | ✓ 867ms | ✓ 987ms | ✓ 916ms | ✓ 1071ms | 否 | http |
| 18.192.100.176:42970 | ✓ 1051ms | 否 | ✓ 1729ms | 否 | ✓ 1817ms | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1295ms | ✓ 1123ms | ✓ 853ms | http |
| 38.145.218.102:8447 | ✓ 1063ms | ✓ 1114ms | ✓ 305ms | ✓ 773ms | ✓ 871ms | http |
| 38.145.220.11:8447 | ✓ 1064ms | ✓ 1109ms | ✓ 310ms | ✓ 821ms | ✓ 830ms | http |
| 38.145.220.198:8448 | ✓ 819ms | ✓ 1109ms | ✓ 551ms | ✓ 890ms | ✓ 631ms | http |
| 38.34.179.65:8450 | 否 | 否 | ✓ 94ms | ✓ 690ms | ✓ 698ms | http |
| 160.250.4.245:1 | ✓ 1553ms | 否 | ✓ 1246ms | ✓ 1312ms | ✓ 1015ms | http |
| 38.145.218.210:8451 | ✓ 103ms | ✓ 712ms | ✓ 266ms | ✓ 703ms | ✓ 534ms | http |
| 38.34.179.27:8452 | ✓ 179ms | ✓ 783ms | ✓ 116ms | ✓ 662ms | ✓ 622ms | http |
| 38.34.179.152:8446 | ✓ 87ms | ✓ 645ms | ✓ 346ms | ✓ 717ms | ✓ 691ms | http |
| 38.34.179.77:8445 | ✓ 114ms | ✓ 686ms | ✓ 275ms | ✓ 810ms | ✓ 612ms | http |
| 38.34.179.174:8453 | ✓ 158ms | ✓ 854ms | ✓ 226ms | ✓ 726ms | ✓ 722ms | http |
| 38.145.203.19:8447 | ✓ 228ms | ✓ 654ms | ✓ 359ms | ✓ 684ms | ✓ 741ms | http |
| 38.34.179.70:8453 | ✓ 93ms | ✓ 688ms | ✓ 294ms | ✓ 690ms | ✓ 617ms | http |
| 38.34.179.179:8449 | ✓ 262ms | ✓ 671ms | ✓ 473ms | ✓ 663ms | ✓ 700ms | http |
| 38.34.179.100:8449 | ✓ 298ms | ✓ 612ms | ✓ 491ms | ✓ 992ms | ✓ 545ms | http |
| 38.34.179.193:8452 | ✓ 199ms | ✓ 689ms | ✓ 517ms | ✓ 657ms | ✓ 832ms | http |
| 45.136.130.169:8446 | ✓ 328ms | ✓ 620ms | ✓ 614ms | ✓ 688ms | ✓ 714ms | http |
| 45.136.130.248:8452 | ✓ 306ms | ✓ 595ms | ✓ 665ms | ✓ 712ms | ✓ 916ms | http |
| 45.136.131.30:8451 | ✓ 268ms | ✓ 930ms | ✓ 400ms | ✓ 1286ms | ✓ 1142ms | http |

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
