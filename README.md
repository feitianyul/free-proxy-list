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

最后更新：2026-04-04 11:36:29 UTC（2026-04-04 19:36:29 UTC+8）

**代理总数：257**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 257 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 257 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 927ms | ✓ 1488ms | ✓ 1017ms | ✓ 1186ms | ✓ 969ms | http |
| 45.136.130.195:8446 | ✓ 1568ms | 否 | ✓ 1203ms | ✓ 1716ms | ✓ 954ms | http |
| 147.161.210.140:8800 | ✓ 1882ms | 否 | ✓ 1053ms | ✓ 1182ms | ✓ 1175ms | http |
| 95.213.217.168:52004 | ✓ 962ms | ✓ 1612ms | ✓ 1665ms | 否 | ✓ 1587ms | http |
| 1.231.81.166:3128 | ✓ 1947ms | ✓ 1237ms | ✓ 1670ms | ✓ 1533ms | ✓ 999ms | http |
| 167.103.34.108:8800 | ✓ 1582ms | 否 | ✓ 1560ms | ✓ 1650ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1541ms | 否 | ✓ 1183ms | 否 | ✓ 1426ms | http |
| 113.160.132.26:8080 | ✓ 1941ms | ✓ 1557ms | ✓ 1773ms | ✓ 1850ms | ✓ 1352ms | http |
| 159.223.71.162:8080 | ✓ 1773ms | 否 | ✓ 1447ms | 否 | ✓ 1348ms | http |
| 159.223.71.162:443 | 否 | 否 | ✓ 1683ms | ✓ 1988ms | ✓ 1541ms | http |
| 45.167.124.52:8080 | ✓ 1273ms | 否 | ✓ 1265ms | ✓ 1861ms | 否 | http |
| 45.136.130.188:8449 | ✓ 269ms | ✓ 1296ms | ✓ 301ms | ✓ 918ms | ✓ 733ms | http |
| 45.136.130.184:8447 | ✓ 355ms | ✓ 1009ms | ✓ 358ms | ✓ 1093ms | ✓ 764ms | http |
| 45.136.131.45:8449 | ✓ 464ms | ✓ 1102ms | ✓ 302ms | ✓ 1057ms | ✓ 854ms | http |
| 45.136.131.44:8449 | ✓ 514ms | ✓ 1094ms | ✓ 575ms | ✓ 959ms | ✓ 737ms | http |
| 45.136.131.40:8449 | ✓ 505ms | ✓ 1599ms | ✓ 336ms | ✓ 985ms | ✓ 758ms | http |
| 45.136.130.247:8448 | ✓ 409ms | ✓ 1365ms | ✓ 925ms | ✓ 1098ms | ✓ 814ms | http |
| 38.34.183.47:8452 | ✓ 662ms | ✓ 928ms | ✓ 722ms | ✓ 1274ms | ✓ 849ms | http |
| 38.145.218.13:8446 | ✓ 404ms | ✓ 1295ms | ✓ 985ms | ✓ 1121ms | ✓ 800ms | http |
| 45.136.130.181:8445 | ✓ 380ms | ✓ 1025ms | ✓ 321ms | ✓ 1104ms | ✓ 791ms | http |
| 38.145.220.35:8446 | ✓ 774ms | ✓ 986ms | ✓ 898ms | ✓ 1072ms | ✓ 821ms | http |
| 45.136.130.177:8447 | ✓ 395ms | 否 | ✓ 350ms | ✓ 1007ms | ✓ 1082ms | http |
| 38.145.208.179:8450 | ✓ 571ms | ✓ 1686ms | ✓ 529ms | ✓ 1215ms | ✓ 875ms | http |
| 38.34.179.61:8445 | ✓ 622ms | ✓ 1147ms | ✓ 562ms | ✓ 1361ms | ✓ 1156ms | http |
| 38.145.208.180:8451 | ✓ 646ms | ✓ 1151ms | ✓ 954ms | ✓ 1274ms | ✓ 860ms | http |
| 38.34.179.24:8447 | ✓ 439ms | 否 | ✓ 894ms | ✓ 987ms | ✓ 736ms | http |
| 38.145.218.235:8453 | ✓ 630ms | 否 | ✓ 391ms | ✓ 1133ms | ✓ 795ms | http |
| 38.145.218.210:8448 | ✓ 425ms | ✓ 1839ms | ✓ 537ms | ✓ 1120ms | ✓ 1055ms | http |
| 38.145.220.100:8451 | ✓ 364ms | ✓ 1599ms | ✓ 573ms | ✓ 1282ms | ✓ 1081ms | http |
| 38.145.203.76:8446 | ✓ 485ms | ✓ 1797ms | ✓ 578ms | ✓ 1122ms | ✓ 1172ms | http |
| 38.34.179.79:8451 | ✓ 928ms | ✓ 1205ms | ✓ 525ms | ✓ 1406ms | ✓ 1080ms | http |
| 38.145.220.33:8446 | ✓ 719ms | ✓ 982ms | ✓ 886ms | ✓ 1163ms | ✓ 1457ms | http |
| 38.34.179.165:8450 | ✓ 728ms | ✓ 1085ms | ✓ 619ms | ✓ 1449ms | ✓ 1033ms | http |
| 38.34.179.177:8446 | ✓ 685ms | ✓ 1019ms | ✓ 515ms | ✓ 1461ms | ✓ 1236ms | http |
| 38.34.183.130:8452 | ✓ 526ms | ✓ 1114ms | ✓ 587ms | ✓ 1823ms | ✓ 760ms | http |
| 38.34.179.49:8445 | ✓ 657ms | ✓ 1189ms | ✓ 1062ms | ✓ 1152ms | ✓ 1158ms | http |
| 45.136.130.187:8449 | ✓ 978ms | ✓ 1607ms | ✓ 339ms | ✓ 980ms | ✓ 726ms | http |
| 45.136.131.32:8446 | ✓ 503ms | ✓ 1090ms | ✓ 708ms | ✓ 1158ms | ✓ 1029ms | http |
| 45.136.131.27:8446 | ✓ 420ms | ✓ 1089ms | ✓ 904ms | ✓ 1041ms | ✓ 886ms | http |
| 45.136.131.26:8446 | ✓ 476ms | 否 | ✓ 428ms | ✓ 1003ms | ✓ 754ms | http |
| 45.136.130.246:8446 | ✓ 557ms | ✓ 1169ms | ✓ 1353ms | ✓ 1288ms | ✓ 776ms | http |
| 38.34.179.154:8453 | ✓ 1276ms | ✓ 1623ms | ✓ 1000ms | ✓ 1193ms | ✓ 1164ms | http |
| 38.145.220.173:8444 | ✓ 1036ms | ✓ 933ms | ✓ 291ms | ✓ 1089ms | ✓ 983ms | http |
| 38.145.208.242:8451 | ✓ 568ms | ✓ 1144ms | ✓ 972ms | ✓ 1161ms | ✓ 952ms | http |
| 38.34.179.95:8444 | ✓ 650ms | ✓ 1140ms | ✓ 673ms | ✓ 1651ms | ✓ 1320ms | http |
| 38.34.179.179:8449 | ✓ 1691ms | ✓ 1076ms | ✓ 567ms | ✓ 1349ms | ✓ 1140ms | http |
| 45.136.130.176:8451 | ✓ 343ms | ✓ 1153ms | ✓ 1829ms | ✓ 1444ms | ✓ 1187ms | http |
| 38.145.220.60:8447 | ✓ 1213ms | ✓ 847ms | ✓ 322ms | ✓ 1095ms | ✓ 1115ms | http |
| 38.34.179.60:8450 | ✓ 1177ms | ✓ 1717ms | ✓ 1314ms | ✓ 1501ms | ✓ 1166ms | http |
| 38.145.220.179:8444 | ✓ 940ms | ✓ 1799ms | ✓ 935ms | ✓ 1248ms | ✓ 1445ms | http |
| 38.34.183.164:8444 | ✓ 596ms | ✓ 1069ms | 否 | ✓ 1702ms | 否 | http |
| 38.145.208.207:8445 | ✓ 429ms | 否 | ✓ 1578ms | ✓ 1135ms | ✓ 867ms | http |
| 38.145.220.41:8444 | ✓ 761ms | 否 | ✓ 725ms | ✓ 1620ms | ✓ 938ms | http |
| 38.145.220.81:8453 | ✓ 511ms | ✓ 1028ms | ✓ 1319ms | 否 | ✓ 1120ms | http |
| 38.145.220.65:8446 | ✓ 532ms | 否 | ✓ 1160ms | ✓ 1508ms | ✓ 1093ms | http |
| 38.145.220.77:8453 | ✓ 548ms | ✓ 1719ms | ✓ 1510ms | ✓ 1433ms | ✓ 977ms | http |
| 45.136.130.182:8446 | 否 | 否 | ✓ 1578ms | ✓ 975ms | ✓ 742ms | http |
| 167.103.144.127:8800 | ✓ 1717ms | 否 | 否 | ✓ 1725ms | ✓ 1671ms | http |
| 38.34.179.186:8444 | 否 | ✓ 1934ms | ✓ 683ms | ✓ 1330ms | ✓ 1261ms | http |
| 217.76.245.80:999 | ✓ 784ms | 否 | ✓ 1246ms | ✓ 1538ms | ✓ 1539ms | http |
| 45.136.131.36:8450 | ✓ 380ms | ✓ 955ms | ✓ 301ms | ✓ 1004ms | ✓ 750ms | http |
| 167.103.31.122:8800 | ✓ 1778ms | 否 | ✓ 1237ms | ✓ 1490ms | 否 | http |
| 198.59.68.130:3128 | ✓ 737ms | ✓ 1971ms | ✓ 1572ms | ✓ 1348ms | ✓ 1327ms | http |
| 35.225.22.61:80 | ✓ 243ms | 否 | ✓ 476ms | ✓ 1032ms | 否 | http |
| 223.16.170.103:80 | ✓ 1402ms | 否 | ✓ 1574ms | 否 | ✓ 1424ms | http |
| 37.204.248.33:443 | ✓ 1322ms | 否 | ✓ 1389ms | ✓ 1660ms | ✓ 1442ms | http |
| 45.136.131.33:8452 | ✓ 688ms | ✓ 1098ms | ✓ 446ms | 否 | ✓ 790ms | http |
| 45.136.131.30:8451 | ✓ 675ms | ✓ 1152ms | ✓ 550ms | 否 | ✓ 745ms | http |
| 38.145.208.215:8444 | ✓ 695ms | ✓ 1077ms | ✓ 611ms | 否 | ✓ 900ms | http |
| 38.145.208.245:8452 | ✓ 683ms | ✓ 961ms | ✓ 746ms | 否 | ✓ 741ms | http |
| 38.145.208.247:8452 | ✓ 642ms | ✓ 1034ms | ✓ 675ms | 否 | ✓ 773ms | http |
| 45.136.130.168:8448 | ✓ 690ms | ✓ 933ms | ✓ 759ms | 否 | ✓ 810ms | http |
| 38.145.208.174:8444 | ✓ 1222ms | 否 | ✓ 693ms | ✓ 1322ms | 否 | http |
| 38.145.208.173:8444 | ✓ 1227ms | 否 | ✓ 735ms | ✓ 1531ms | 否 | http |
| 38.145.208.177:8450 | ✓ 1204ms | 否 | ✓ 720ms | ✓ 1292ms | 否 | http |
| 38.145.208.246:8450 | ✓ 658ms | 否 | ✓ 368ms | ✓ 955ms | ✓ 770ms | http |
| 38.145.208.170:8453 | ✓ 1208ms | 否 | ✓ 726ms | ✓ 1280ms | 否 | http |
| 45.136.131.67:8453 | ✓ 1428ms | 否 | ✓ 1486ms | ✓ 1063ms | ✓ 1157ms | http |
| 45.136.131.59:8444 | ✓ 1560ms | 否 | ✓ 1463ms | ✓ 1054ms | ✓ 1148ms | http |
| 45.136.130.171:8450 | ✓ 671ms | ✓ 952ms | ✓ 730ms | 否 | ✓ 1125ms | http |
| 45.136.130.175:8450 | ✓ 642ms | ✓ 975ms | ✓ 742ms | 否 | ✓ 1153ms | http |
| 111.227.254.9:22222 | ✓ 1253ms | 否 | 否 | ✓ 1874ms | ✓ 1299ms | http |
| 38.34.179.172:8447 | ✓ 1072ms | ✓ 1247ms | ✓ 1439ms | ✓ 1295ms | ✓ 957ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1303ms | ✓ 1109ms | ✓ 1329ms | ✓ 1076ms | http |
| 111.227.254.12:22222 | ✓ 1138ms | ✓ 1600ms | ✓ 1607ms | ✓ 1892ms | ✓ 1225ms | http |
| 45.167.125.21:999 | ✓ 569ms | 否 | ✓ 1166ms | ✓ 1985ms | ✓ 1613ms | http |
| 38.145.208.241:8447 | ✓ 562ms | ✓ 1593ms | ✓ 1141ms | ✓ 1047ms | ✓ 802ms | http |
| 38.145.218.161:8445 | ✓ 1232ms | ✓ 1398ms | ✓ 605ms | ✓ 1042ms | ✓ 1340ms | http |
| 45.136.131.31:8451 | ✓ 689ms | 否 | ✓ 925ms | ✓ 978ms | ✓ 898ms | http |
| 209.38.154.7:1080 | ✓ 1168ms | 否 | 否 | ✓ 920ms | ✓ 738ms | http |
| 38.145.208.222:8443 | ✓ 699ms | 否 | ✓ 705ms | ✓ 1237ms | ✓ 968ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 1240ms | ✓ 1618ms | ✓ 1143ms | http |
| 38.34.183.8:8450 | ✓ 1745ms | ✓ 1610ms | ✓ 1318ms | 否 | ✓ 1269ms | http |
| 38.34.179.151:8446 | ✓ 403ms | ✓ 1295ms | ✓ 343ms | ✓ 1049ms | ✓ 779ms | http |
| 192.71.213.85:9090 | ✓ 1171ms | 否 | ✓ 1933ms | ✓ 1816ms | 否 | http |
| 38.145.208.253:8443 | ✓ 842ms | ✓ 1241ms | ✓ 405ms | ✓ 1019ms | ✓ 909ms | http |
| 38.34.179.106:8445 | ✓ 809ms | ✓ 1127ms | ✓ 533ms | ✓ 994ms | ✓ 823ms | http |
| 38.34.179.80:8452 | ✓ 1077ms | 否 | ✓ 777ms | ✓ 1309ms | ✓ 1279ms | http |
| 38.34.179.175:8445 | ✓ 1093ms | 否 | ✓ 660ms | ✓ 1610ms | ✓ 1336ms | http |
| 38.34.179.103:8448 | 否 | 否 | ✓ 1822ms | ✓ 1236ms | ✓ 882ms | http |

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
