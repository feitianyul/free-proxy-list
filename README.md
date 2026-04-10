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

最后更新：2026-04-10 03:53:08 UTC（2026-04-10 11:53:08 UTC+8）

**代理总数：543**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 543 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 543 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.209.239.31:30000 | ✓ 1709ms | ✓ 825ms | ✓ 916ms | 否 | ✓ 1019ms | http |
| 38.34.179.67:8446 | ✓ 1884ms | 否 | ✓ 585ms | ✓ 747ms | ✓ 1157ms | http |
| 147.161.210.140:8800 | ✓ 1708ms | 否 | ✓ 540ms | ✓ 1165ms | ✓ 1151ms | http |
| 167.103.115.102:8800 | ✓ 1633ms | 否 | ✓ 965ms | ✓ 1031ms | ✓ 1124ms | http |
| 5.104.87.17:8051 | ✓ 1192ms | 否 | ✓ 1488ms | 否 | ✓ 849ms | http |
| 38.145.208.204:8446 | ✓ 1885ms | 否 | ✓ 1233ms | ✓ 681ms | ✓ 760ms | http |
| 147.161.239.240:8800 | ✓ 1213ms | ✓ 1848ms | ✓ 1393ms | ✓ 1668ms | ✓ 1428ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1331ms | ✓ 1170ms | ✓ 1581ms | ✓ 1011ms | http |
| 152.32.132.190:7890 | ✓ 951ms | 否 | 否 | ✓ 1837ms | ✓ 853ms | http |
| 212.113.102.190:33362 | ✓ 1238ms | 否 | ✓ 1228ms | ✓ 1899ms | ✓ 1629ms | http |
| 167.103.34.108:8800 | ✓ 1678ms | 否 | ✓ 1446ms | ✓ 1803ms | ✓ 1716ms | http |
| 38.145.208.174:8444 | ✓ 1871ms | 否 | ✓ 1365ms | ✓ 1012ms | ✓ 1482ms | http |
| 38.145.208.173:8444 | ✓ 1871ms | 否 | ✓ 1361ms | ✓ 1014ms | ✓ 1497ms | http |
| 38.92.10.98:20058 | ✓ 202ms | ✓ 835ms | ✓ 1108ms | ✓ 897ms | ✓ 545ms | http |
| 38.145.208.224:8445 | ✓ 946ms | ✓ 783ms | ✓ 395ms | 否 | ✓ 1203ms | http |
| 167.103.144.127:8800 | ✓ 1136ms | ✓ 1722ms | ✓ 1171ms | ✓ 1631ms | ✓ 1556ms | http |
| 45.136.131.40:8444 | ✓ 250ms | 否 | ✓ 978ms | ✓ 847ms | ✓ 1110ms | http |
| 85.239.59.252:7890 | 否 | ✓ 1947ms | ✓ 786ms | ✓ 1963ms | 否 | http |
| 103.82.23.118:5224 | ✓ 1992ms | 否 | ✓ 1620ms | ✓ 1812ms | 否 | http |
| 185.76.240.21:10001 | ✓ 1036ms | 否 | ✓ 1075ms | 否 | ✓ 1539ms | http |
| 38.34.179.49:8444 | ✓ 527ms | ✓ 1420ms | ✓ 861ms | 否 | 否 | http |
| 1.231.81.166:3128 | ✓ 1280ms | 否 | ✓ 1452ms | ✓ 940ms | 否 | http |
| 185.76.241.184:10001 | ✓ 1141ms | 否 | ✓ 1930ms | 否 | ✓ 1748ms | http |
| 45.167.124.52:8080 | ✓ 718ms | ✓ 1667ms | ✓ 885ms | 否 | 否 | http |
| 185.76.240.29:10001 | ✓ 1049ms | 否 | ✓ 1091ms | 否 | ✓ 1572ms | http |
| 185.76.240.184:10001 | ✓ 924ms | 否 | ✓ 1074ms | 否 | ✓ 1593ms | http |
| 185.76.241.110:10001 | ✓ 1104ms | 否 | ✓ 1053ms | 否 | ✓ 1683ms | http |
| 185.76.240.226:10001 | ✓ 1006ms | 否 | ✓ 1023ms | 否 | ✓ 1794ms | http |
| 45.88.0.113:3128 | ✓ 1431ms | 否 | ✓ 1124ms | 否 | ✓ 1227ms | http |
| 167.103.31.122:8800 | ✓ 1304ms | 否 | ✓ 1270ms | ✓ 1962ms | 否 | http |
| 54.37.72.89:80 | ✓ 1310ms | 否 | ✓ 1175ms | 否 | ✓ 1955ms | http |
| 38.147.160.208:24239 | ✓ 482ms | ✓ 788ms | ✓ 976ms | ✓ 909ms | ✓ 546ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 453ms | ✓ 1191ms | ✓ 1094ms | http |
| 59.46.216.131:30001 | ✓ 921ms | 否 | ✓ 1000ms | 否 | ✓ 1035ms | http |
| 181.78.49.177:999 | ✓ 1935ms | ✓ 1727ms | ✓ 788ms | ✓ 1490ms | ✓ 1308ms | http |
| 38.145.208.209:8447 | ✓ 633ms | ✓ 920ms | ✓ 289ms | ✓ 883ms | ✓ 517ms | http |
| 38.145.208.234:8446 | ✓ 680ms | ✓ 1588ms | ✓ 115ms | ✓ 665ms | ✓ 521ms | http |
| 38.145.220.39:8449 | ✓ 650ms | 否 | ✓ 165ms | ✓ 805ms | ✓ 907ms | http |
| 38.34.179.87:8447 | ✓ 1339ms | ✓ 910ms | ✓ 331ms | 否 | 否 | http |
| 45.136.131.25:8453 | ✓ 880ms | ✓ 987ms | ✓ 495ms | ✓ 1089ms | ✓ 596ms | http |
| 38.145.203.96:8451 | ✓ 877ms | ✓ 1984ms | ✓ 143ms | 否 | 否 | http |
| 45.136.131.37:8447 | ✓ 640ms | 否 | ✓ 901ms | ✓ 1506ms | 否 | http |
| 38.145.203.35:8450 | ✓ 1092ms | ✓ 1136ms | ✓ 144ms | ✓ 807ms | ✓ 1858ms | http |
| 38.34.179.82:8445 | ✓ 662ms | 否 | ✓ 104ms | ✓ 701ms | ✓ 569ms | http |
| 38.145.208.229:8453 | ✓ 675ms | 否 | ✓ 91ms | ✓ 679ms | ✓ 559ms | http |
| 45.136.131.43:8447 | ✓ 656ms | ✓ 1295ms | ✓ 1590ms | 否 | 否 | http |
| 38.145.203.41:8453 | ✓ 1130ms | ✓ 1124ms | ✓ 313ms | ✓ 1049ms | 否 | http |
| 38.34.179.213:8443 | ✓ 646ms | ✓ 1598ms | ✓ 1390ms | 否 | 否 | http |
| 45.136.131.33:8452 | ✓ 868ms | ✓ 1009ms | ✓ 675ms | ✓ 1141ms | 否 | http |
| 38.34.179.11:8447 | ✓ 880ms | ✓ 1757ms | ✓ 244ms | ✓ 824ms | 否 | http |
| 38.145.203.32:8452 | ✓ 1555ms | ✓ 901ms | ✓ 337ms | ✓ 1024ms | 否 | http |
| 38.145.220.77:8453 | ✓ 885ms | ✓ 1750ms | ✓ 518ms | ✓ 691ms | 否 | http |
| 45.136.130.178:8453 | ✓ 1102ms | 否 | ✓ 111ms | ✓ 693ms | 否 | http |
| 38.34.179.199:8452 | ✓ 653ms | ✓ 1599ms | ✓ 460ms | ✓ 1231ms | 否 | http |
| 38.34.179.46:8448 | ✓ 655ms | ✓ 854ms | ✓ 925ms | ✓ 899ms | ✓ 519ms | http |
| 45.136.131.54:8448 | ✓ 875ms | 否 | ✓ 98ms | ✓ 829ms | 否 | http |
| 38.145.203.106:8447 | ✓ 893ms | 否 | ✓ 310ms | ✓ 700ms | 否 | http |
| 38.34.179.29:8452 | ✓ 623ms | ✓ 990ms | ✓ 679ms | ✓ 1530ms | 否 | http |
| 38.34.179.21:8452 | ✓ 643ms | ✓ 976ms | ✓ 694ms | ✓ 1515ms | 否 | http |
| 166.1.18.178:7890 | 否 | ✓ 1074ms | ✓ 947ms | 否 | ✓ 1260ms | http |
| 38.34.179.228:8453 | ✓ 419ms | ✓ 811ms | ✓ 581ms | ✓ 1453ms | ✓ 530ms | http |
| 38.145.208.215:8444 | ✓ 497ms | ✓ 1019ms | ✓ 301ms | ✓ 1036ms | ✓ 847ms | http |
| 38.145.208.220:8448 | ✓ 708ms | ✓ 863ms | ✓ 246ms | ✓ 1028ms | ✓ 1152ms | http |
| 45.136.131.31:8451 | ✓ 413ms | ✓ 1024ms | ✓ 454ms | ✓ 1515ms | ✓ 1318ms | http |
| 38.34.179.194:8451 | ✓ 422ms | ✓ 1547ms | ✓ 1313ms | ✓ 863ms | ✓ 625ms | http |
| 101.43.127.100:8877 | ✓ 840ms | ✓ 1637ms | ✓ 805ms | ✓ 1158ms | ✓ 1867ms | http |
| 194.163.183.242:3128 | ✓ 1174ms | ✓ 1687ms | ✓ 1917ms | ✓ 1985ms | ✓ 1521ms | http |
| 45.12.151.226:2829 | ✓ 1441ms | ✓ 1960ms | ✓ 1698ms | 否 | ✓ 1954ms | http |
| 77.91.77.220:3128 | ✓ 1159ms | 否 | ✓ 998ms | 否 | ✓ 1639ms | http |
| 186.96.111.214:999 | ✓ 1199ms | ✓ 1874ms | ✓ 1076ms | ✓ 1638ms | ✓ 1380ms | http |
| 45.167.125.21:999 | ✓ 1290ms | ✓ 1715ms | ✓ 1504ms | 否 | ✓ 1667ms | http |
| 177.234.217.88:999 | ✓ 1659ms | 否 | ✓ 1724ms | 否 | ✓ 1912ms | http |
| 45.136.131.28:8449 | ✓ 87ms | ✓ 817ms | ✓ 71ms | ✓ 662ms | ✓ 515ms | http |
| 45.136.131.27:8444 | ✓ 82ms | ✓ 823ms | ✓ 85ms | ✓ 683ms | ✓ 513ms | http |
| 38.145.203.97:8446 | ✓ 180ms | ✓ 807ms | ✓ 195ms | ✓ 861ms | ✓ 544ms | http |
| 38.145.218.229:8444 | ✓ 272ms | ✓ 1897ms | ✓ 77ms | ✓ 753ms | ✓ 534ms | http |
| 185.76.241.157:10001 | ✓ 752ms | 否 | ✓ 886ms | 否 | ✓ 1884ms | http |
| 185.76.240.201:10001 | ✓ 999ms | 否 | ✓ 763ms | 否 | ✓ 1905ms | http |
| 185.76.240.51:10001 | ✓ 1003ms | 否 | ✓ 764ms | 否 | ✓ 1932ms | http |
| 38.145.208.182:8452 | ✓ 437ms | ✓ 835ms | ✓ 767ms | ✓ 968ms | ✓ 1365ms | http |
| 45.136.131.36:8450 | ✓ 417ms | ✓ 1018ms | ✓ 439ms | ✓ 1404ms | ✓ 839ms | http |
| 45.88.0.115:3128 | ✓ 1225ms | ✓ 1535ms | ✓ 1818ms | ✓ 1564ms | ✓ 1766ms | http |
| 213.220.62.63:3128 | ✓ 1228ms | 否 | ✓ 1343ms | 否 | ✓ 1671ms | http |
| 213.220.62.62:3128 | 否 | 否 | ✓ 669ms | ✓ 1591ms | ✓ 1701ms | http |
| 45.88.0.111:3128 | ✓ 1227ms | ✓ 1544ms | 否 | 否 | ✓ 1198ms | http |
| 185.76.240.64:10001 | ✓ 1231ms | 否 | ✓ 1092ms | 否 | ✓ 1724ms | http |
| 185.76.240.236:10001 | ✓ 1255ms | 否 | ✓ 1086ms | 否 | ✓ 1610ms | http |
| 185.76.240.203:10001 | ✓ 1228ms | 否 | ✓ 1825ms | 否 | ✓ 1664ms | http |
| 185.76.240.61:10001 | ✓ 1231ms | 否 | ✓ 1641ms | 否 | ✓ 1643ms | http |
| 38.34.179.17:8445 | ✓ 131ms | ✓ 887ms | ✓ 183ms | ✓ 687ms | ✓ 579ms | http |
| 38.34.179.76:8444 | ✓ 160ms | ✓ 903ms | ✓ 142ms | ✓ 792ms | ✓ 534ms | http |
| 38.34.183.47:8451 | ✓ 176ms | ✓ 812ms | ✓ 565ms | ✓ 1243ms | ✓ 497ms | http |
| 45.136.131.58:8445 | ✓ 109ms | ✓ 852ms | ✓ 100ms | ✓ 895ms | ✓ 574ms | http |
| 38.145.218.161:8451 | ✓ 221ms | ✓ 891ms | ✓ 555ms | ✓ 1217ms | ✓ 492ms | http |
| 38.145.208.181:8445 | ✓ 106ms | ✓ 820ms | ✓ 150ms | ✓ 699ms | ✓ 813ms | http |
| 45.136.130.190:8453 | ✓ 545ms | ✓ 785ms | ✓ 153ms | ✓ 780ms | ✓ 1027ms | http |
| 38.145.208.221:8444 | ✓ 541ms | ✓ 1325ms | ✓ 91ms | ✓ 732ms | ✓ 760ms | http |
| 38.34.179.154:8444 | ✓ 351ms | ✓ 1480ms | ✓ 641ms | ✓ 1761ms | ✓ 549ms | http |
| 38.34.179.91:8447 | ✓ 802ms | 否 | ✓ 246ms | ✓ 882ms | ✓ 1605ms | http |
| 38.34.179.94:8451 | ✓ 885ms | ✓ 1456ms | ✓ 105ms | ✓ 657ms | ✓ 1209ms | http |

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
