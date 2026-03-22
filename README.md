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

最后更新：2026-03-22 06:50:26 UTC（2026-03-22 14:50:26 UTC+8）

**代理总数：170**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 170 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 170 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 685ms | ✓ 1312ms | ✓ 812ms | ✓ 1598ms | ✓ 872ms | http |
| 219.117.204.211:7799 | ✓ 589ms | 否 | 否 | ✓ 841ms | ✓ 664ms | http |
| 1.231.81.166:3128 | ✓ 750ms | ✓ 1246ms | ✓ 1488ms | ✓ 1297ms | ✓ 1023ms | http |
| 147.161.239.240:8800 | ✓ 1280ms | ✓ 1878ms | ✓ 1263ms | ✓ 1505ms | ✓ 1323ms | http |
| 167.103.34.108:8800 | ✓ 1859ms | 否 | ✓ 1428ms | ✓ 1538ms | 否 | http |
| 121.126.185.63:25152 | ✓ 1354ms | 否 | ✓ 1355ms | ✓ 1703ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1082ms | ✓ 1919ms | ✓ 1582ms | ✓ 1959ms | ✓ 1638ms | http |
| 47.77.193.180:1080 | ✓ 723ms | ✓ 1474ms | ✓ 360ms | ✓ 664ms | ✓ 496ms | http |
| 38.34.179.25:8444 | ✓ 688ms | ✓ 1812ms | ✓ 512ms | ✓ 810ms | ✓ 587ms | http |
| 38.34.179.27:8451 | ✓ 710ms | ✓ 1710ms | ✓ 804ms | ✓ 837ms | ✓ 548ms | http |
| 43.99.54.236:5555 | ✓ 606ms | ✓ 1185ms | ✓ 612ms | ✓ 775ms | ✓ 617ms | http |
| 38.34.179.105:8449 | ✓ 1245ms | ✓ 1365ms | ✓ 841ms | ✓ 821ms | ✓ 596ms | http |
| 38.34.179.173:8452 | ✓ 1737ms | ✓ 1720ms | ✓ 287ms | ✓ 903ms | ✓ 1537ms | http |
| 38.34.179.203:8451 | ✓ 710ms | ✓ 1935ms | ✓ 769ms | ✓ 861ms | 否 | http |
| 104.168.158.236:10808 | ✓ 1139ms | 否 | ✓ 1650ms | ✓ 851ms | ✓ 641ms | http |
| 113.160.132.26:8080 | ✓ 833ms | ✓ 1269ms | ✓ 914ms | ✓ 1169ms | ✓ 890ms | http |
| 45.149.92.147:5001 | ✓ 895ms | 否 | ✓ 1391ms | ✓ 1682ms | ✓ 727ms | http |
| 150.241.77.172:1080 | ✓ 673ms | 否 | ✓ 1277ms | ✓ 1635ms | ✓ 1542ms | http |
| 106.75.15.167:7890 | 否 | ✓ 1134ms | 否 | ✓ 1461ms | ✓ 1363ms | http |
| 144.31.79.117:8888 | ✓ 1121ms | 否 | ✓ 1428ms | 否 | ✓ 1547ms | http |
| 212.192.12.90:6005 | ✓ 1712ms | 否 | ✓ 1368ms | ✓ 1344ms | 否 | http |
| 38.34.179.83:8448 | ✓ 1183ms | 否 | ✓ 1554ms | 否 | ✓ 1837ms | http |
| 38.34.179.150:8449 | ✓ 918ms | 否 | ✓ 1218ms | ✓ 1080ms | 否 | http |
| 38.34.183.222:8453 | ✓ 1752ms | ✓ 644ms | ✓ 582ms | ✓ 1428ms | ✓ 1810ms | http |
| 167.103.31.122:8800 | ✓ 1332ms | 否 | ✓ 1298ms | 否 | ✓ 1492ms | http |
| 91.238.105.64:2024 | ✓ 1295ms | 否 | ✓ 1967ms | 否 | ✓ 1935ms | http |
| 38.34.179.98:8453 | ✓ 853ms | ✓ 786ms | ✓ 192ms | 否 | 否 | http |
| 45.136.131.55:8444 | ✓ 93ms | ✓ 938ms | ✓ 602ms | ✓ 892ms | ✓ 583ms | http |
| 45.136.131.39:8444 | ✓ 189ms | ✓ 773ms | ✓ 164ms | ✓ 839ms | ✓ 596ms | http |
| 45.136.130.167:8446 | ✓ 149ms | 否 | ✓ 144ms | ✓ 977ms | ✓ 536ms | http |
| 38.34.179.86:8452 | ✓ 932ms | 否 | ✓ 117ms | ✓ 781ms | ✓ 692ms | http |
| 38.34.179.14:8444 | ✓ 941ms | 否 | ✓ 172ms | ✓ 807ms | ✓ 558ms | http |
| 38.145.203.39:8444 | 否 | ✓ 797ms | ✓ 692ms | ✓ 1263ms | ✓ 1891ms | http |
| 38.34.179.164:8444 | ✓ 920ms | 否 | ✓ 862ms | ✓ 720ms | ✓ 538ms | http |
| 38.34.183.224:8445 | ✓ 923ms | 否 | ✓ 641ms | ✓ 705ms | ✓ 1079ms | http |
| 38.145.208.196:8444 | ✓ 938ms | 否 | ✓ 838ms | ✓ 1016ms | ✓ 627ms | http |
| 160.250.5.22:1 | ✓ 1482ms | 否 | ✓ 1056ms | ✓ 1240ms | ✓ 1070ms | http |
| 160.250.4.245:1 | ✓ 1480ms | 否 | ✓ 1058ms | ✓ 1424ms | ✓ 1005ms | http |
| 38.34.179.60:8450 | ✓ 1481ms | ✓ 1715ms | ✓ 534ms | 否 | ✓ 999ms | http |
| 38.145.208.185:8449 | ✓ 1109ms | 否 | ✓ 890ms | ✓ 1838ms | ✓ 787ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1043ms | ✓ 1265ms | ✓ 1990ms | http |
| 34.158.226.196:3128 | ✓ 1899ms | 否 | ✓ 1741ms | 否 | ✓ 1770ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1467ms | ✓ 1874ms | 否 | ✓ 1180ms | http |
| 38.34.179.61:8452 | ✓ 879ms | ✓ 758ms | ✓ 275ms | ✓ 951ms | ✓ 583ms | http |
| 38.145.203.97:8444 | ✓ 837ms | ✓ 609ms | ✓ 311ms | ✓ 828ms | ✓ 535ms | http |
| 45.136.130.183:8450 | ✓ 836ms | ✓ 658ms | ✓ 266ms | ✓ 699ms | ✓ 507ms | http |
| 38.34.179.178:8444 | ✓ 853ms | ✓ 729ms | ✓ 176ms | ✓ 715ms | ✓ 537ms | http |
| 38.34.179.189:8452 | ✓ 855ms | ✓ 1872ms | ✓ 189ms | ✓ 1037ms | ✓ 928ms | http |
| 38.145.218.87:8444 | ✓ 837ms | ✓ 1002ms | ✓ 811ms | ✓ 1243ms | ✓ 573ms | http |
| 8.219.97.248:80 | ✓ 1908ms | 否 | 否 | ✓ 1189ms | ✓ 1832ms | http |
| 38.145.203.19:8444 | ✓ 879ms | ✓ 1827ms | ✓ 1082ms | ✓ 697ms | ✓ 660ms | http |
| 38.145.203.43:8445 | 否 | 否 | ✓ 90ms | ✓ 691ms | ✓ 1018ms | http |
| 38.34.179.73:8444 | ✓ 879ms | ✓ 790ms | ✓ 1920ms | ✓ 1019ms | ✓ 677ms | http |
| 2.56.173.45:10808 | 否 | ✓ 1606ms | ✓ 1352ms | 否 | ✓ 1709ms | http |
| 59.46.216.131:30001 | ✓ 1114ms | 否 | 否 | ✓ 1297ms | ✓ 1068ms | http |
| 181.78.44.63:999 | ✓ 1004ms | ✓ 1995ms | ✓ 932ms | ✓ 1941ms | 否 | http |
| 104.129.202.127:10810 | ✓ 401ms | ✓ 685ms | ✓ 601ms | ✓ 758ms | ✓ 523ms | http |
| 104.129.202.127:12354 | ✓ 534ms | ✓ 680ms | ✓ 576ms | ✓ 719ms | ✓ 493ms | http |
| 137.184.1.87:3128 | ✓ 1278ms | ✓ 762ms | ✓ 729ms | ✓ 641ms | ✓ 490ms | http |
| 38.34.179.172:8451 | ✓ 374ms | ✓ 1389ms | ✓ 110ms | ✓ 700ms | ✓ 528ms | http |
| 38.34.178.141:8453 | ✓ 414ms | ✓ 1475ms | ✓ 397ms | ✓ 692ms | ✓ 522ms | http |
| 103.84.95.54:7890 | ✓ 632ms | 否 | ✓ 602ms | 否 | ✓ 639ms | http |
| 38.34.178.7:8452 | ✓ 399ms | ✓ 683ms | ✓ 947ms | ✓ 825ms | ✓ 593ms | http |
| 38.34.178.241:8453 | ✓ 420ms | ✓ 1933ms | ✓ 103ms | ✓ 671ms | ✓ 503ms | http |
| 38.34.178.175:8452 | ✓ 403ms | ✓ 1247ms | ✓ 588ms | ✓ 665ms | ✓ 557ms | http |
| 38.34.179.78:8445 | ✓ 684ms | ✓ 684ms | ✓ 276ms | ✓ 1562ms | ✓ 1792ms | http |
| 103.52.114.95:3128 | ✓ 1446ms | 否 | ✓ 946ms | ✓ 1204ms | ✓ 904ms | http |
| 38.145.203.19:8449 | ✓ 411ms | ✓ 734ms | ✓ 272ms | ✓ 1228ms | ✓ 1772ms | http |
| 72.56.79.129:1080 | 否 | 否 | ✓ 1561ms | ✓ 1496ms | ✓ 1130ms | http |
| 103.139.138.194:3128 | ✓ 1768ms | 否 | ✓ 1058ms | ✓ 1414ms | ✓ 1126ms | http |
| 38.34.179.54:8446 | ✓ 1708ms | 否 | ✓ 364ms | 否 | ✓ 1507ms | http |
| 38.34.179.61:8445 | ✓ 404ms | ✓ 1194ms | ✓ 1466ms | ✓ 954ms | ✓ 630ms | http |
| 45.136.131.54:8448 | ✓ 529ms | ✓ 1409ms | 否 | ✓ 1839ms | 否 | http |
| 38.34.179.165:8446 | ✓ 932ms | ✓ 1826ms | ✓ 345ms | ✓ 763ms | ✓ 874ms | http |
| 38.34.179.23:8444 | ✓ 795ms | ✓ 1837ms | ✓ 191ms | ✓ 926ms | ✓ 1222ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 947ms | ✓ 1351ms | ✓ 969ms | http |
| 38.34.179.162:8451 | ✓ 294ms | ✓ 1196ms | 否 | ✓ 992ms | ✓ 640ms | http |
| 137.220.150.22:6005 | ✓ 1062ms | 否 | 否 | ✓ 1965ms | ✓ 1557ms | http |
| 38.34.179.178:8445 | ✓ 1438ms | ✓ 1633ms | ✓ 274ms | ✓ 867ms | ✓ 1754ms | http |
| 45.136.131.35:8452 | ✓ 637ms | 否 | ✓ 1071ms | ✓ 798ms | ✓ 700ms | http |
| 137.220.150.104:6005 | ✓ 1492ms | 否 | ✓ 873ms | ✓ 1292ms | ✓ 979ms | http |
| 38.34.179.26:8450 | ✓ 636ms | ✓ 1070ms | ✓ 1342ms | ✓ 787ms | ✓ 855ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1164ms | ✓ 1109ms | ✓ 1874ms | ✓ 1885ms | http |
| 194.67.99.223:1080 | ✓ 1613ms | 否 | ✓ 1677ms | ✓ 1964ms | ✓ 1530ms | http |
| 120.92.212.16:7890 | ✓ 908ms | ✓ 1140ms | ✓ 1525ms | 否 | 否 | http |
| 38.34.183.13:8449 | ✓ 636ms | ✓ 774ms | ✓ 731ms | ✓ 1996ms | 否 | http |
| 45.136.130.162:8443 | ✓ 501ms | ✓ 1236ms | ✓ 85ms | ✓ 655ms | ✓ 493ms | http |
| 38.34.178.245:8446 | ✓ 892ms | ✓ 650ms | ✓ 110ms | ✓ 682ms | ✓ 625ms | http |
| 150.107.140.238:3128 | ✓ 745ms | 否 | ✓ 961ms | ✓ 1075ms | ✓ 852ms | http |
| 38.34.179.155:8452 | ✓ 693ms | ✓ 704ms | ✓ 253ms | ✓ 1136ms | 否 | http |
| 38.34.179.174:8453 | ✓ 1560ms | 否 | ✓ 314ms | ✓ 1667ms | 否 | http |
| 36.155.100.217:8080 | ✓ 1220ms | 否 | ✓ 937ms | ✓ 1522ms | ✓ 1178ms | http |
| 38.145.208.179:8451 | 否 | 否 | ✓ 253ms | ✓ 1274ms | ✓ 930ms | http |
| 115.231.181.40:8128 | ✓ 985ms | ✓ 1190ms | 否 | 否 | ✓ 1524ms | http |
| 38.145.208.242:8451 | 否 | 否 | ✓ 911ms | ✓ 668ms | ✓ 506ms | http |
| 38.34.178.153:8453 | 否 | 否 | ✓ 886ms | ✓ 968ms | ✓ 1089ms | http |
| 77.232.135.22:1080 | ✓ 743ms | 否 | ✓ 1775ms | ✓ 1963ms | ✓ 1366ms | http |
| 77.110.113.24:40000 | ✓ 732ms | ✓ 1928ms | ✓ 1907ms | 否 | ✓ 1939ms | http |
| 150.249.255.91:3128 | ✓ 1459ms | ✓ 1112ms | ✓ 490ms | 否 | 否 | http |
| 112.163.160.93:3128 | 否 | 否 | ✓ 832ms | ✓ 1084ms | ✓ 752ms | http |

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
