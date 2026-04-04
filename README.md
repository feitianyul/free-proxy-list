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

最后更新：2026-04-04 22:33:40 UTC（2026-04-05 06:33:40 UTC+8）

**代理总数：166**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 166 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 166 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 821ms | ✓ 1066ms | ✓ 1300ms | ✓ 1113ms | ✓ 909ms | http |
| 1.231.81.166:3128 | ✓ 765ms | ✓ 882ms | ✓ 1711ms | ✓ 1362ms | ✓ 953ms | http |
| 45.136.131.36:8450 | ✓ 569ms | ✓ 1418ms | ✓ 1306ms | ✓ 787ms | ✓ 1014ms | http |
| 45.136.131.43:8447 | ✓ 644ms | ✓ 1373ms | ✓ 1278ms | ✓ 821ms | ✓ 920ms | http |
| 111.227.254.9:22222 | ✓ 930ms | ✓ 1300ms | ✓ 1021ms | ✓ 1333ms | ✓ 1008ms | http |
| 111.227.254.12:22222 | ✓ 1020ms | ✓ 1292ms | ✓ 935ms | ✓ 1285ms | ✓ 1102ms | http |
| 167.103.115.102:8800 | ✓ 1378ms | 否 | ✓ 985ms | ✓ 1076ms | ✓ 1354ms | http |
| 113.160.132.26:8080 | ✓ 1665ms | ✓ 1621ms | ✓ 1313ms | ✓ 1164ms | ✓ 1009ms | http |
| 159.223.71.162:443 | ✓ 1385ms | 否 | 否 | ✓ 1010ms | ✓ 819ms | http |
| 159.223.71.162:8080 | ✓ 1379ms | 否 | 否 | ✓ 1050ms | ✓ 834ms | http |
| 167.103.34.108:8800 | ✓ 1444ms | 否 | ✓ 1421ms | ✓ 1508ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1070ms | 否 | 否 | ✓ 1805ms | ✓ 1971ms | http |
| 209.38.154.7:1080 | ✓ 1083ms | 否 | 否 | ✓ 680ms | ✓ 1321ms | http |
| 180.250.219.58:53281 | ✓ 1633ms | 否 | ✓ 1530ms | ✓ 1636ms | ✓ 1908ms | http |
| 38.145.218.208:8446 | ✓ 615ms | ✓ 702ms | ✓ 433ms | ✓ 713ms | ✓ 561ms | http |
| 38.145.220.13:8450 | ✓ 634ms | ✓ 682ms | ✓ 458ms | ✓ 784ms | ✓ 529ms | http |
| 38.34.179.179:8449 | ✓ 569ms | ✓ 713ms | ✓ 430ms | ✓ 814ms | ✓ 562ms | http |
| 38.34.179.192:8446 | ✓ 558ms | ✓ 713ms | ✓ 430ms | ✓ 776ms | ✓ 578ms | http |
| 38.34.183.130:8452 | ✓ 671ms | ✓ 633ms | ✓ 505ms | ✓ 791ms | ✓ 527ms | http |
| 38.145.220.182:8450 | ✓ 662ms | ✓ 977ms | ✓ 250ms | ✓ 784ms | ✓ 892ms | http |
| 38.145.220.188:8450 | ✓ 625ms | ✓ 971ms | ✓ 281ms | ✓ 798ms | ✓ 858ms | http |
| 45.136.130.195:8446 | ✓ 525ms | ✓ 716ms | ✓ 478ms | ✓ 1794ms | ✓ 544ms | http |
| 45.136.130.194:8451 | ✓ 547ms | ✓ 731ms | ✓ 478ms | ✓ 1373ms | ✓ 869ms | http |
| 38.34.183.13:8449 | ✓ 654ms | ✓ 701ms | ✓ 454ms | ✓ 878ms | ✓ 1541ms | http |
| 168.110.52.228:3128 | ✓ 993ms | ✓ 998ms | ✓ 1434ms | ✓ 738ms | ✓ 571ms | http |
| 38.34.179.6:8449 | ✓ 1002ms | ✓ 1056ms | ✓ 328ms | ✓ 1056ms | ✓ 789ms | http |
| 45.136.130.169:8444 | ✓ 1162ms | ✓ 873ms | ✓ 95ms | ✓ 796ms | ✓ 1228ms | http |
| 45.136.131.39:8451 | ✓ 674ms | ✓ 756ms | ✓ 633ms | ✓ 1419ms | ✓ 552ms | http |
| 167.103.144.127:8800 | ✓ 1155ms | 否 | ✓ 1260ms | ✓ 1207ms | ✓ 1064ms | http |
| 218.108.131.186:17890 | ✓ 1120ms | ✓ 1150ms | 否 | ✓ 1692ms | ✓ 843ms | http |
| 160.238.65.7:3128 | ✓ 703ms | ✓ 1520ms | ✓ 1693ms | 否 | ✓ 1439ms | http |
| 38.145.218.234:8447 | 否 | 否 | ✓ 1972ms | ✓ 1318ms | ✓ 1748ms | http |
| 167.103.31.122:8800 | ✓ 1885ms | 否 | ✓ 1666ms | 否 | ✓ 1518ms | http |
| 45.136.131.37:8447 | ✓ 546ms | ✓ 1144ms | ✓ 1108ms | ✓ 1054ms | ✓ 1187ms | http |
| 38.145.208.244:8446 | ✓ 1112ms | ✓ 1379ms | ✓ 616ms | ✓ 1530ms | ✓ 1617ms | http |
| 45.136.131.47:8452 | ✓ 1317ms | ✓ 797ms | ✓ 366ms | ✓ 1676ms | ✓ 1527ms | http |
| 38.145.220.72:8445 | ✓ 1607ms | 否 | ✓ 1530ms | 否 | ✓ 1953ms | http |
| 95.213.217.168:52004 | ✓ 1160ms | 否 | ✓ 1570ms | 否 | ✓ 1649ms | http |
| 45.167.124.52:8080 | ✓ 761ms | ✓ 1700ms | ✓ 886ms | ✓ 1788ms | ✓ 1442ms | http |
| 138.197.68.35:4857 | ✓ 1556ms | 否 | ✓ 504ms | ✓ 1530ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1026ms | ✓ 1712ms | ✓ 801ms | ✓ 1220ms | 否 | http |
| 160.238.65.6:3128 | ✓ 644ms | ✓ 1526ms | ✓ 607ms | ✓ 1555ms | 否 | http |
| 160.238.65.4:3128 | ✓ 639ms | ✓ 1631ms | ✓ 620ms | 否 | 否 | http |
| 116.80.65.78:3172 | ✓ 1472ms | 否 | ✓ 1453ms | 否 | ✓ 1619ms | http |
| 120.92.212.16:7890 | ✓ 1005ms | 否 | ✓ 1231ms | ✓ 1177ms | ✓ 938ms | http |
| 38.145.220.11:8447 | ✓ 548ms | ✓ 939ms | 否 | ✓ 1064ms | ✓ 1556ms | http |
| 38.145.208.242:8444 | ✓ 714ms | ✓ 770ms | ✓ 925ms | ✓ 788ms | ✓ 625ms | http |
| 101.43.127.100:8877 | ✓ 797ms | ✓ 1051ms | ✓ 1304ms | ✓ 1101ms | ✓ 912ms | http |
| 147.161.239.240:8800 | ✓ 1299ms | ✓ 1884ms | ✓ 1303ms | ✓ 1431ms | ✓ 1284ms | http |
| 120.92.212.16:8890 | ✓ 1154ms | ✓ 1175ms | ✓ 1583ms | ✓ 1174ms | ✓ 956ms | http |
| 91.233.223.147:3128 | ✓ 1365ms | 否 | ✓ 1781ms | 否 | ✓ 1859ms | http |
| 177.234.217.88:999 | ✓ 1221ms | 否 | ✓ 1914ms | 否 | ✓ 1717ms | http |
| 45.167.125.21:999 | ✓ 921ms | ✓ 1860ms | ✓ 1437ms | 否 | ✓ 1749ms | http |
| 38.145.203.98:8447 | ✓ 1904ms | 否 | ✓ 1771ms | ✓ 1754ms | 否 | http |
| 45.136.131.46:8443 | ✓ 393ms | ✓ 975ms | ✓ 172ms | ✓ 677ms | ✓ 527ms | http |
| 38.145.220.168:8444 | ✓ 1139ms | ✓ 795ms | ✓ 338ms | ✓ 799ms | ✓ 927ms | http |
| 38.145.203.87:8444 | ✓ 1113ms | ✓ 813ms | ✓ 353ms | ✓ 816ms | ✓ 955ms | http |
| 38.145.220.102:8453 | ✓ 613ms | ✓ 1119ms | ✓ 532ms | ✓ 827ms | ✓ 1077ms | http |
| 38.145.218.210:8448 | ✓ 821ms | ✓ 945ms | ✓ 497ms | ✓ 884ms | ✓ 857ms | http |
| 38.34.183.233:8445 | ✓ 628ms | ✓ 698ms | ✓ 118ms | ✓ 799ms | ✓ 844ms | http |
| 38.145.220.103:8446 | ✓ 1646ms | ✓ 1910ms | ✓ 1662ms | ✓ 1132ms | ✓ 1402ms | http |
| 45.136.130.176:8451 | 否 | ✓ 856ms | ✓ 833ms | ✓ 1728ms | ✓ 1175ms | http |
| 38.145.220.40:8445 | ✓ 152ms | ✓ 826ms | ✓ 593ms | ✓ 955ms | ✓ 706ms | http |
| 38.34.179.25:8448 | ✓ 552ms | ✓ 814ms | ✓ 216ms | ✓ 832ms | ✓ 580ms | http |
| 45.136.130.181:8445 | ✓ 521ms | ✓ 859ms | ✓ 196ms | ✓ 1043ms | ✓ 891ms | http |
| 38.145.208.244:8444 | ✓ 489ms | ✓ 876ms | ✓ 317ms | ✓ 1023ms | ✓ 1313ms | http |
| 38.145.208.253:8444 | ✓ 501ms | ✓ 922ms | ✓ 324ms | ✓ 984ms | ✓ 1289ms | http |
| 59.46.216.131:30001 | ✓ 930ms | ✓ 1299ms | ✓ 1089ms | ✓ 1288ms | ✓ 991ms | http |
| 45.136.131.28:8449 | ✓ 617ms | ✓ 785ms | ✓ 111ms | ✓ 721ms | ✓ 612ms | http |
| 154.64.230.89:3128 | ✓ 516ms | ✓ 714ms | ✓ 885ms | ✓ 832ms | ✓ 541ms | http |
| 45.136.130.180:8452 | ✓ 619ms | ✓ 941ms | ✓ 283ms | ✓ 957ms | ✓ 642ms | http |
| 38.145.220.39:8452 | ✓ 551ms | ✓ 692ms | ✓ 470ms | ✓ 1465ms | ✓ 693ms | http |
| 38.145.220.179:8444 | ✓ 595ms | ✓ 1153ms | ✓ 229ms | ✓ 932ms | ✓ 923ms | http |
| 38.145.218.212:8448 | ✓ 967ms | ✓ 965ms | ✓ 110ms | ✓ 789ms | ✓ 686ms | http |
| 38.145.218.232:8448 | ✓ 427ms | ✓ 617ms | ✓ 98ms | ✓ 720ms | ✓ 975ms | http |
| 38.145.220.173:8444 | ✓ 415ms | ✓ 948ms | ✓ 267ms | ✓ 1193ms | ✓ 1102ms | http |
| 45.136.130.187:8449 | 否 | 否 | ✓ 1056ms | ✓ 1721ms | ✓ 929ms | http |
| 38.34.179.29:8452 | ✓ 1125ms | ✓ 1572ms | ✓ 324ms | ✓ 1011ms | ✓ 1194ms | http |
| 45.136.130.192:8450 | ✓ 510ms | ✓ 1012ms | ✓ 1300ms | ✓ 1585ms | ✓ 685ms | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1755ms | ✓ 1490ms | ✓ 1514ms | http |
| 116.80.81.13:7777 | 否 | 否 | ✓ 1477ms | ✓ 1758ms | ✓ 1654ms | http |
| 116.80.63.67:7777 | ✓ 1628ms | ✓ 1924ms | 否 | ✓ 1740ms | 否 | http |
| 38.145.203.76:8446 | ✓ 1750ms | ✓ 1657ms | ✓ 1166ms | 否 | ✓ 1908ms | http |
| 45.136.130.166:8452 | ✓ 577ms | ✓ 939ms | ✓ 791ms | ✓ 849ms | ✓ 830ms | http |
| 45.136.131.25:8453 | ✓ 566ms | ✓ 1152ms | ✓ 585ms | ✓ 1081ms | ✓ 736ms | http |
| 38.34.183.164:8444 | ✓ 388ms | ✓ 739ms | ✓ 116ms | ✓ 791ms | ✓ 568ms | http |
| 38.145.203.124:8444 | ✓ 478ms | ✓ 931ms | ✓ 606ms | ✓ 815ms | ✓ 1206ms | http |
| 35.225.22.61:80 | ✓ 1044ms | 否 | ✓ 992ms | ✓ 1188ms | 否 | http |
| 38.145.203.86:8449 | ✓ 750ms | ✓ 1669ms | ✓ 938ms | ✓ 1895ms | 否 | http |
| 38.145.220.82:8448 | ✓ 230ms | ✓ 804ms | ✓ 153ms | ✓ 828ms | ✓ 651ms | http |
| 38.145.220.27:8445 | ✓ 630ms | ✓ 1321ms | ✓ 156ms | ✓ 904ms | ✓ 578ms | http |
| 38.34.178.152:8451 | ✓ 877ms | ✓ 1087ms | ✓ 361ms | ✓ 1331ms | ✓ 1556ms | http |
| 38.145.218.134:8446 | ✓ 884ms | ✓ 1196ms | ✓ 284ms | ✓ 860ms | ✓ 590ms | http |
| 38.145.218.163:8446 | ✓ 881ms | ✓ 1235ms | ✓ 273ms | ✓ 832ms | ✓ 578ms | http |
| 38.34.183.11:8446 | ✓ 887ms | ✓ 1219ms | ✓ 272ms | ✓ 847ms | ✓ 559ms | http |
| 38.34.179.189:8445 | ✓ 869ms | ✓ 698ms | ✓ 760ms | ✓ 1299ms | ✓ 999ms | http |
| 38.145.218.161:8445 | ✓ 892ms | ✓ 1363ms | ✓ 273ms | ✓ 805ms | ✓ 625ms | http |
| 38.145.220.93:8445 | ✓ 880ms | ✓ 975ms | ✓ 708ms | ✓ 1883ms | ✓ 728ms | http |
| 38.145.220.60:8447 | ✓ 778ms | ✓ 889ms | ✓ 188ms | ✓ 938ms | ✓ 750ms | http |
| 43.167.237.94:3128 | 否 | 否 | ✓ 892ms | ✓ 847ms | ✓ 1304ms | http |

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
