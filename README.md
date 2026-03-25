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

最后更新：2026-03-25 19:54:46 UTC（2026-03-26 03:54:46 UTC+8）

**代理总数：481**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 481 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 481 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.208.242:8451 | ✓ 394ms | ✓ 673ms | ✓ 106ms | ✓ 767ms | ✓ 659ms | http |
| 38.145.208.244:8448 | ✓ 800ms | ✓ 698ms | ✓ 107ms | ✓ 684ms | ✓ 505ms | http |
| 38.34.179.105:8449 | ✓ 394ms | ✓ 1199ms | ✓ 486ms | ✓ 668ms | ✓ 645ms | http |
| 38.34.179.161:8448 | ✓ 394ms | ✓ 616ms | ✓ 1017ms | ✓ 1467ms | ✓ 518ms | http |
| 38.34.179.162:8451 | ✓ 394ms | ✓ 612ms | ✓ 1014ms | ✓ 1500ms | ✓ 502ms | http |
| 38.34.179.150:8449 | ✓ 394ms | ✓ 1235ms | ✓ 1383ms | ✓ 672ms | ✓ 524ms | http |
| 38.34.183.13:8449 | ✓ 381ms | ✓ 600ms | ✓ 574ms | 否 | ✓ 1005ms | http |
| 38.145.208.241:8453 | ✓ 802ms | ✓ 622ms | ✓ 116ms | ✓ 790ms | ✓ 572ms | http |
| 38.34.179.6:8449 | ✓ 800ms | ✓ 642ms | ✓ 122ms | ✓ 700ms | ✓ 918ms | http |
| 45.136.131.53:8452 | ✓ 918ms | ✓ 636ms | ✓ 303ms | ✓ 1115ms | 否 | http |
| 35.225.22.61:80 | ✓ 598ms | ✓ 1543ms | ✓ 939ms | ✓ 1384ms | ✓ 1147ms | http |
| 115.231.181.40:8128 | ✓ 905ms | ✓ 1002ms | ✓ 926ms | ✓ 1044ms | ✓ 938ms | http |
| 219.117.204.211:7799 | ✓ 1233ms | 否 | ✓ 958ms | ✓ 917ms | ✓ 688ms | http |
| 38.145.208.185:8449 | ✓ 536ms | ✓ 757ms | ✓ 1416ms | ✓ 924ms | ✓ 583ms | http |
| 147.161.210.140:8800 | ✓ 1228ms | 否 | ✓ 959ms | ✓ 1079ms | ✓ 791ms | http |
| 38.34.179.27:8451 | ✓ 379ms | ✓ 1000ms | 否 | ✓ 694ms | ✓ 739ms | http |
| 167.103.115.102:8800 | ✓ 897ms | 否 | ✓ 921ms | ✓ 1016ms | ✓ 1019ms | http |
| 38.34.179.186:8444 | ✓ 921ms | ✓ 1032ms | ✓ 1037ms | 否 | ✓ 525ms | http |
| 45.136.130.193:8444 | ✓ 377ms | ✓ 794ms | ✓ 1392ms | ✓ 1478ms | ✓ 1135ms | http |
| 45.136.130.197:8444 | ✓ 377ms | ✓ 819ms | ✓ 1443ms | ✓ 1398ms | ✓ 1146ms | http |
| 222.228.171.92:8080 | ✓ 1255ms | ✓ 1865ms | ✓ 1392ms | ✓ 1100ms | ✓ 1026ms | http |
| 137.220.150.104:6005 | ✓ 1187ms | 否 | ✓ 909ms | ✓ 1182ms | ✓ 896ms | http |
| 113.160.132.26:8080 | ✓ 1358ms | ✓ 1486ms | ✓ 908ms | ✓ 1326ms | ✓ 1004ms | http |
| 103.84.95.54:7890 | ✓ 1342ms | ✓ 1666ms | ✓ 694ms | ✓ 1245ms | ✓ 1673ms | http |
| 38.34.179.86:8452 | ✓ 515ms | ✓ 1102ms | ✓ 1555ms | ✓ 1735ms | ✓ 680ms | http |
| 45.136.130.189:8451 | ✓ 377ms | ✓ 1071ms | ✓ 1917ms | ✓ 1616ms | ✓ 1849ms | http |
| 38.34.179.16:8451 | ✓ 1110ms | ✓ 616ms | ✓ 1758ms | 否 | ✓ 536ms | http |
| 45.136.130.188:8449 | ✓ 377ms | ✓ 1282ms | ✓ 1723ms | ✓ 1350ms | ✓ 1505ms | http |
| 167.103.34.108:8800 | ✓ 1264ms | ✓ 1906ms | 否 | ✓ 1408ms | ✓ 1318ms | http |
| 34.101.184.164:3128 | ✓ 864ms | 否 | ✓ 1393ms | ✓ 1945ms | ✓ 1256ms | http |
| 47.77.193.180:1080 | ✓ 175ms | ✓ 739ms | ✓ 231ms | ✓ 730ms | ✓ 500ms | http |
| 38.34.179.173:8452 | ✓ 183ms | ✓ 645ms | ✓ 263ms | ✓ 761ms | ✓ 886ms | http |
| 38.34.179.40:8446 | ✓ 270ms | ✓ 820ms | ✓ 148ms | ✓ 1026ms | ✓ 516ms | http |
| 38.34.179.78:8445 | ✓ 192ms | ✓ 784ms | ✓ 632ms | ✓ 710ms | ✓ 572ms | http |
| 38.34.179.51:8449 | ✓ 148ms | ✓ 790ms | ✓ 124ms | ✓ 1003ms | ✓ 515ms | http |
| 38.145.218.229:8450 | ✓ 443ms | ✓ 1527ms | ✓ 126ms | ✓ 694ms | ✓ 725ms | http |
| 38.145.220.198:8448 | ✓ 198ms | ✓ 829ms | ✓ 925ms | ✓ 971ms | ✓ 722ms | http |
| 38.34.179.98:8453 | ✓ 280ms | ✓ 984ms | ✓ 785ms | ✓ 853ms | ✓ 763ms | http |
| 38.34.179.96:8451 | ✓ 295ms | ✓ 992ms | ✓ 770ms | ✓ 848ms | ✓ 590ms | http |
| 38.145.220.33:8448 | ✓ 232ms | ✓ 815ms | ✓ 793ms | ✓ 1166ms | ✓ 970ms | http |
| 38.34.179.178:8445 | ✓ 169ms | ✓ 648ms | ✓ 936ms | ✓ 1254ms | ✓ 520ms | http |
| 45.136.131.35:8452 | ✓ 197ms | ✓ 695ms | ✓ 456ms | 否 | ✓ 1304ms | http |
| 38.34.179.174:8453 | ✓ 164ms | ✓ 649ms | ✓ 571ms | ✓ 1509ms | ✓ 525ms | http |
| 175.194.173.105:3128 | ✓ 651ms | ✓ 1256ms | ✓ 964ms | ✓ 934ms | ✓ 745ms | http |
| 112.163.160.93:3128 | ✓ 833ms | ✓ 1383ms | ✓ 807ms | ✓ 923ms | ✓ 987ms | http |
| 210.223.44.230:3128 | ✓ 669ms | ✓ 976ms | ✓ 1036ms | ✓ 1727ms | ✓ 671ms | http |
| 38.34.179.177:8447 | ✓ 250ms | ✓ 925ms | ✓ 1212ms | ✓ 777ms | ✓ 522ms | http |
| 38.34.183.233:8448 | ✓ 1698ms | ✓ 603ms | ✓ 594ms | 否 | ✓ 1321ms | http |
| 45.136.131.68:8444 | ✓ 987ms | 否 | ✓ 565ms | ✓ 718ms | ✓ 1378ms | http |
| 45.136.131.64:8444 | ✓ 644ms | 否 | ✓ 989ms | ✓ 728ms | ✓ 1366ms | http |
| 45.136.131.67:8444 | ✓ 644ms | 否 | ✓ 988ms | ✓ 690ms | ✓ 1416ms | http |
| 38.34.183.232:8453 | ✓ 1656ms | ✓ 602ms | ✓ 452ms | ✓ 953ms | 否 | http |
| 38.145.218.228:8447 | ✓ 194ms | ✓ 862ms | ✓ 884ms | ✓ 1776ms | ✓ 570ms | http |
| 120.92.212.16:8890 | ✓ 917ms | ✓ 1152ms | ✓ 928ms | ✓ 1178ms | ✓ 931ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1374ms | ✓ 1153ms | ✓ 1399ms | ✓ 1154ms | http |
| 20.210.39.153:8561 | ✓ 609ms | ✓ 919ms | ✓ 600ms | ✓ 968ms | ✓ 938ms | http |
| 101.47.73.135:3128 | ✓ 1711ms | 否 | 否 | ✓ 1754ms | ✓ 982ms | http |
| 20.78.118.91:8561 | ✓ 625ms | ✓ 1022ms | ✓ 631ms | ✓ 966ms | ✓ 898ms | http |
| 38.34.183.234:8450 | ✓ 1716ms | ✓ 604ms | ✓ 471ms | ✓ 922ms | 否 | http |
| 20.78.26.206:8561 | ✓ 622ms | ✓ 1063ms | ✓ 664ms | ✓ 929ms | ✓ 928ms | http |
| 202.141.161.53:10808 | ✓ 938ms | ✓ 1279ms | ✓ 1442ms | ✓ 1176ms | 否 | http |
| 64.227.76.27:1080 | ✓ 1142ms | 否 | ✓ 1517ms | 否 | ✓ 1990ms | http |
| 167.103.31.122:8800 | ✓ 1506ms | 否 | ✓ 1643ms | ✓ 1735ms | ✓ 1691ms | http |
| 137.220.151.110:6005 | ✓ 1804ms | 否 | ✓ 1288ms | ✓ 1494ms | ✓ 1210ms | http |
| 1.231.81.166:3128 | ✓ 1081ms | ✓ 1350ms | ✓ 1682ms | ✓ 1629ms | ✓ 1332ms | http |
| 85.208.51.165:443 | ✓ 1132ms | ✓ 1841ms | ✓ 1782ms | 否 | 否 | http |
| 106.75.15.167:7890 | ✓ 1087ms | ✓ 1094ms | 否 | ✓ 1148ms | ✓ 1126ms | http |
| 45.136.130.191:8453 | ✓ 245ms | ✓ 810ms | ✓ 1375ms | ✓ 1903ms | ✓ 968ms | http |
| 186.148.180.46:999 | ✓ 1161ms | 否 | ✓ 1818ms | ✓ 1779ms | 否 | http |
| 137.184.1.87:3128 | ✓ 292ms | ✓ 726ms | ✓ 548ms | ✓ 624ms | ✓ 490ms | http |
| 137.184.6.37:3128 | ✓ 296ms | ✓ 586ms | ✓ 834ms | ✓ 636ms | ✓ 482ms | http |
| 38.145.208.221:8453 | ✓ 1048ms | ✓ 605ms | ✓ 909ms | ✓ 1022ms | ✓ 1827ms | http |
| 23.95.76.201:8443 | ✓ 667ms | ✓ 1278ms | ✓ 1105ms | ✓ 1411ms | ✓ 1129ms | http |
| 150.107.140.238:3128 | ✓ 1253ms | 否 | ✓ 854ms | 否 | ✓ 889ms | http |
| 116.80.49.156:3172 | ✓ 1833ms | 否 | ✓ 1495ms | ✓ 1770ms | 否 | http |
| 103.217.216.65:8181 | ✓ 1902ms | 否 | ✓ 1511ms | ✓ 1402ms | 否 | http |
| 85.208.108.43:2094 | ✓ 976ms | 否 | ✓ 452ms | ✓ 1113ms | ✓ 807ms | http |
| 43.99.54.236:5555 | ✓ 622ms | ✓ 899ms | ✓ 615ms | ✓ 787ms | ✓ 630ms | http |
| 150.31.45.65:8118 | ✓ 1409ms | 否 | ✓ 1335ms | 否 | ✓ 967ms | http |
| 116.80.65.77:3172 | ✓ 1884ms | 否 | 否 | ✓ 1792ms | ✓ 1798ms | http |
| 38.145.203.45:8443 | ✓ 566ms | ✓ 773ms | ✓ 241ms | ✓ 805ms | ✓ 584ms | http |
| 38.145.208.222:8443 | ✓ 577ms | ✓ 685ms | ✓ 484ms | ✓ 820ms | ✓ 636ms | http |
| 38.145.208.220:8443 | ✓ 556ms | ✓ 702ms | ✓ 471ms | ✓ 834ms | ✓ 633ms | http |
| 38.145.208.216:8443 | ✓ 572ms | ✓ 703ms | ✓ 471ms | ✓ 828ms | ✓ 638ms | http |
| 38.145.208.217:8443 | ✓ 571ms | ✓ 701ms | ✓ 467ms | ✓ 835ms | ✓ 631ms | http |
| 38.34.179.38:8447 | ✓ 573ms | ✓ 704ms | ✓ 605ms | ✓ 810ms | ✓ 509ms | http |
| 38.34.179.97:8448 | ✓ 549ms | ✓ 741ms | ✓ 812ms | ✓ 725ms | ✓ 500ms | http |
| 38.34.179.104:8448 | ✓ 577ms | ✓ 694ms | ✓ 932ms | ✓ 680ms | ✓ 503ms | http |
| 38.145.218.234:8443 | ✓ 587ms | ✓ 753ms | ✓ 247ms | ✓ 953ms | ✓ 887ms | http |
| 38.34.179.61:8445 | ✓ 593ms | ✓ 700ms | ✓ 913ms | ✓ 703ms | ✓ 523ms | http |
| 38.145.208.215:8443 | ✓ 567ms | ✓ 691ms | ✓ 475ms | ✓ 832ms | ✓ 857ms | http |
| 38.145.218.218:8443 | ✓ 578ms | ✓ 730ms | ✓ 269ms | ✓ 938ms | ✓ 911ms | http |
| 38.145.203.132:8443 | ✓ 573ms | ✓ 688ms | ✓ 911ms | ✓ 685ms | ✓ 499ms | http |
| 38.145.208.237:8443 | ✓ 564ms | ✓ 700ms | ✓ 904ms | ✓ 703ms | ✓ 492ms | http |
| 38.145.208.238:8443 | ✓ 555ms | ✓ 681ms | ✓ 919ms | ✓ 725ms | ✓ 500ms | http |
| 38.145.208.230:8443 | ✓ 553ms | ✓ 703ms | ✓ 898ms | ✓ 732ms | ✓ 517ms | http |
| 38.145.208.175:8443 | ✓ 571ms | ✓ 732ms | ✓ 872ms | ✓ 772ms | ✓ 561ms | http |
| 38.145.208.171:8443 | ✓ 588ms | ✓ 748ms | ✓ 854ms | ✓ 688ms | ✓ 651ms | http |
| 38.145.208.170:8443 | ✓ 547ms | ✓ 719ms | ✓ 886ms | ✓ 687ms | ✓ 654ms | http |
| 38.145.208.173:8443 | ✓ 576ms | ✓ 674ms | ✓ 920ms | ✓ 688ms | ✓ 657ms | http |

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
