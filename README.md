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

最后更新：2026-03-21 08:34:09 UTC（2026-03-21 16:34:09 UTC+8）

**代理总数：105**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 105 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 105 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1328ms | 否 | ✓ 846ms | ✓ 846ms | ✓ 839ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1534ms | ✓ 1284ms | ✓ 927ms | http |
| 137.220.150.104:6005 | ✓ 1441ms | 否 | ✓ 703ms | ✓ 1054ms | ✓ 907ms | http |
| 137.220.150.152:6005 | ✓ 1477ms | ✓ 1772ms | ✓ 915ms | ✓ 1563ms | ✓ 827ms | http |
| 219.117.204.211:7799 | ✓ 1324ms | 否 | ✓ 1037ms | 否 | ✓ 946ms | http |
| 147.161.239.240:8800 | ✓ 1485ms | 否 | ✓ 1738ms | ✓ 1943ms | ✓ 1655ms | http |
| 167.103.34.108:8800 | ✓ 1856ms | 否 | ✓ 1553ms | ✓ 1596ms | 否 | http |
| 35.225.22.61:80 | ✓ 766ms | 否 | ✓ 362ms | 否 | ✓ 1141ms | http |
| 101.47.73.135:3128 | ✓ 1569ms | 否 | 否 | ✓ 1831ms | ✓ 1002ms | http |
| 120.92.212.16:7890 | ✓ 1386ms | 否 | ✓ 754ms | ✓ 1392ms | ✓ 1349ms | http |
| 133.242.138.34:8100 | ✓ 1302ms | ✓ 1231ms | ✓ 929ms | ✓ 1875ms | ✓ 1740ms | http |
| 167.103.31.122:8800 | ✓ 1859ms | 否 | ✓ 1673ms | 否 | ✓ 1985ms | http |
| 59.46.216.131:30001 | ✓ 1948ms | 否 | ✓ 870ms | ✓ 1560ms | 否 | http |
| 45.88.0.114:3128 | ✓ 1159ms | ✓ 1981ms | ✓ 1669ms | 否 | ✓ 1718ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1612ms | ✓ 1324ms | ✓ 1091ms | http |
| 45.88.0.117:3128 | 否 | ✓ 1556ms | ✓ 966ms | 否 | ✓ 1644ms | http |
| 38.34.179.78:8445 | ✓ 408ms | ✓ 991ms | ✓ 615ms | ✓ 780ms | ✓ 613ms | http |
| 120.92.212.16:8890 | ✓ 739ms | ✓ 1155ms | ✓ 1573ms | 否 | 否 | http |
| 38.34.179.74:8449 | ✓ 1188ms | ✓ 773ms | 否 | ✓ 1856ms | ✓ 1235ms | http |
| 172.212.68.37:3128 | ✓ 946ms | ✓ 1609ms | ✓ 1407ms | ✓ 1833ms | ✓ 1156ms | http |
| 137.220.150.22:6005 | 否 | 否 | ✓ 1116ms | ✓ 1190ms | ✓ 865ms | http |
| 103.82.23.118:5221 | ✓ 1225ms | 否 | ✓ 915ms | 否 | ✓ 1465ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1023ms | ✓ 1927ms | 否 | ✓ 1251ms | http |
| 38.145.208.172:8448 | ✓ 695ms | ✓ 1477ms | ✓ 779ms | ✓ 787ms | ✓ 1616ms | http |
| 27.49.68.66:9999 | ✓ 1177ms | 否 | 否 | ✓ 1259ms | ✓ 1283ms | http |
| 181.78.44.63:999 | ✓ 981ms | ✓ 1815ms | ✓ 1412ms | 否 | ✓ 1675ms | http |
| 150.249.255.91:3128 | ✓ 1997ms | 否 | 否 | ✓ 1244ms | ✓ 635ms | http |
| 103.51.205.117:8090 | ✓ 1796ms | 否 | 否 | ✓ 1349ms | ✓ 1320ms | http |
| 38.34.179.172:8451 | ✓ 999ms | ✓ 1245ms | ✓ 183ms | ✓ 770ms | ✓ 662ms | http |
| 137.220.151.110:6005 | ✓ 1170ms | 否 | ✓ 727ms | ✓ 1059ms | ✓ 815ms | http |
| 38.34.179.173:8452 | ✓ 996ms | ✓ 1645ms | ✓ 424ms | ✓ 1534ms | 否 | http |
| 210.223.44.230:3128 | ✓ 502ms | 否 | ✓ 501ms | ✓ 805ms | ✓ 623ms | http |
| 91.238.105.64:2024 | ✓ 1098ms | 否 | ✓ 1096ms | ✓ 1766ms | ✓ 1379ms | http |
| 116.80.49.162:3172 | ✓ 1478ms | 否 | 否 | ✓ 1961ms | ✓ 1786ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1319ms | ✓ 1394ms | ✓ 1894ms | 否 | http |
| 47.77.193.180:1080 | ✓ 1885ms | ✓ 1346ms | ✓ 275ms | ✓ 730ms | ✓ 630ms | http |
| 103.113.70.189:1081 | ✓ 535ms | 否 | ✓ 959ms | ✓ 1207ms | ✓ 871ms | http |
| 106.75.15.167:7890 | ✓ 1676ms | ✓ 830ms | ✓ 1572ms | ✓ 1095ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1040ms | ✓ 1410ms | ✓ 1130ms | http |
| 201.150.116.3:999 | 否 | 否 | ✓ 905ms | ✓ 1833ms | ✓ 1188ms | http |
| 217.174.244.117:3129 | ✓ 1896ms | ✓ 1937ms | ✓ 1682ms | 否 | ✓ 1939ms | http |
| 121.230.8.41:1080 | ✓ 1908ms | ✓ 1915ms | ✓ 1674ms | ✓ 1972ms | ✓ 1107ms | http |
| 166.88.55.83:7890 | ✓ 568ms | ✓ 1009ms | ✓ 562ms | ✓ 713ms | ✓ 580ms | http |
| 202.141.161.53:30001 | ✓ 1886ms | ✓ 1664ms | ✓ 1325ms | 否 | ✓ 936ms | http |
| 106.117.208.101:7890 | ✓ 753ms | 否 | ✓ 1846ms | ✓ 1006ms | ✓ 728ms | http |
| 45.149.92.147:5001 | ✓ 667ms | 否 | ✓ 665ms | ✓ 1657ms | ✓ 631ms | http |
| 114.237.77.231:1080 | ✓ 700ms | ✓ 883ms | ✓ 799ms | 否 | ✓ 810ms | http |
| 103.39.51.190:8080 | ✓ 1665ms | 否 | 否 | ✓ 1501ms | ✓ 1695ms | http |
| 47.101.149.27:9010 | ✓ 1213ms | 否 | ✓ 1190ms | 否 | ✓ 1282ms | http |
| 222.109.119.178:3128 | ✓ 1694ms | ✓ 1284ms | ✓ 1578ms | 否 | 否 | http |
| 24.199.124.151:3128 | ✓ 315ms | ✓ 1276ms | ✓ 545ms | ✓ 765ms | ✓ 580ms | http |
| 103.52.114.95:3128 | ✓ 1458ms | 否 | ✓ 816ms | ✓ 1141ms | ✓ 895ms | http |
| 160.250.4.245:1 | ✓ 1592ms | 否 | ✓ 1262ms | ✓ 1150ms | ✓ 934ms | http |
| 148.153.56.51:80 | 否 | ✓ 914ms | ✓ 1750ms | ✓ 1657ms | ✓ 1464ms | http |
| 65.108.203.35:18080 | ✓ 1720ms | 否 | 否 | ✓ 1859ms | ✓ 1834ms | http |
| 45.136.131.25:8450 | ✓ 535ms | ✓ 872ms | ✓ 175ms | ✓ 905ms | ✓ 941ms | http |
| 165.227.236.67:3128 | ✓ 656ms | 否 | ✓ 1675ms | ✓ 1828ms | 否 | http |
| 38.34.183.234:8450 | ✓ 287ms | ✓ 1125ms | ✓ 747ms | ✓ 1282ms | ✓ 591ms | http |
| 38.34.179.186:8444 | ✓ 930ms | ✓ 821ms | ✓ 407ms | ✓ 831ms | ✓ 638ms | http |
| 38.145.203.132:8452 | ✓ 728ms | ✓ 1342ms | ✓ 181ms | ✓ 897ms | ✓ 718ms | http |
| 45.136.130.174:8450 | ✓ 1093ms | 否 | ✓ 1892ms | ✓ 758ms | ✓ 840ms | http |
| 38.145.208.215:8453 | ✓ 547ms | ✓ 945ms | ✓ 179ms | ✓ 893ms | ✓ 683ms | http |
| 38.34.179.105:8444 | ✓ 682ms | ✓ 901ms | ✓ 170ms | ✓ 823ms | ✓ 627ms | http |
| 38.34.179.88:8445 | ✓ 545ms | ✓ 836ms | ✓ 255ms | ✓ 991ms | ✓ 769ms | http |
| 38.145.208.177:8451 | ✓ 567ms | ✓ 1336ms | ✓ 279ms | ✓ 771ms | ✓ 566ms | http |
| 38.34.179.179:8449 | ✓ 544ms | ✓ 766ms | ✓ 377ms | ✓ 992ms | ✓ 644ms | http |
| 38.145.220.65:8452 | ✓ 703ms | ✓ 1194ms | ✓ 284ms | ✓ 812ms | ✓ 624ms | http |
| 38.34.179.160:8445 | ✓ 564ms | ✓ 1798ms | ✓ 182ms | ✓ 845ms | ✓ 622ms | http |
| 38.145.208.193:8452 | ✓ 566ms | ✓ 1925ms | ✓ 165ms | ✓ 805ms | ✓ 590ms | http |
| 38.145.220.41:8452 | ✓ 703ms | 否 | ✓ 203ms | ✓ 804ms | ✓ 670ms | http |
| 38.145.203.34:8444 | ✓ 566ms | ✓ 728ms | ✓ 1692ms | ✓ 810ms | ✓ 691ms | http |
| 38.34.179.84:8445 | ✓ 544ms | ✓ 1812ms | ✓ 200ms | ✓ 880ms | ✓ 611ms | http |
| 38.145.220.49:8444 | ✓ 545ms | ✓ 1465ms | ✓ 324ms | ✓ 1093ms | ✓ 764ms | http |
| 38.145.220.38:8452 | ✓ 566ms | ✓ 1829ms | ✓ 871ms | ✓ 801ms | ✓ 612ms | http |
| 38.34.179.83:8449 | ✓ 544ms | 否 | ✓ 197ms | ✓ 893ms | ✓ 593ms | http |
| 38.145.220.96:8453 | ✓ 1222ms | ✓ 806ms | ✓ 735ms | ✓ 1183ms | ✓ 1170ms | http |
| 38.145.208.224:8445 | ✓ 562ms | ✓ 966ms | 否 | ✓ 775ms | ✓ 620ms | http |
| 38.34.179.48:8448 | ✓ 703ms | 否 | ✓ 185ms | ✓ 795ms | ✓ 700ms | http |
| 38.145.203.108:8445 | ✓ 838ms | ✓ 1809ms | ✓ 482ms | ✓ 1115ms | ✓ 1630ms | http |
| 38.145.220.14:8449 | ✓ 838ms | 否 | ✓ 430ms | ✓ 783ms | ✓ 623ms | http |
| 38.34.179.167:8450 | ✓ 1565ms | ✓ 1451ms | ✓ 1039ms | ✓ 1595ms | ✓ 727ms | http |
| 38.145.208.171:8449 | ✓ 1248ms | ✓ 1268ms | ✓ 419ms | ✓ 1187ms | ✓ 597ms | http |
| 38.34.179.181:8446 | ✓ 1092ms | ✓ 1933ms | ✓ 524ms | ✓ 1232ms | ✓ 653ms | http |
| 45.136.130.168:8448 | ✓ 278ms | ✓ 740ms | ✓ 197ms | ✓ 966ms | ✓ 715ms | http |
| 45.136.130.251:8452 | ✓ 365ms | ✓ 1589ms | ✓ 379ms | ✓ 817ms | 否 | http |
| 45.136.130.170:8448 | ✓ 255ms | ✓ 1173ms | ✓ 190ms | ✓ 898ms | ✓ 617ms | http |
| 45.136.131.54:8450 | ✓ 352ms | ✓ 1665ms | ✓ 838ms | 否 | 否 | http |
| 45.136.130.173:8448 | ✓ 238ms | ✓ 1269ms | ✓ 225ms | ✓ 969ms | ✓ 623ms | http |
| 38.145.208.194:8453 | ✓ 1836ms | ✓ 1554ms | 否 | ✓ 1361ms | ✓ 1204ms | http |
| 38.145.203.32:8452 | ✓ 566ms | ✓ 770ms | ✓ 1653ms | ✓ 1331ms | ✓ 629ms | http |
| 45.136.130.252:8448 | ✓ 430ms | ✓ 1507ms | ✓ 216ms | ✓ 822ms | ✓ 683ms | http |
| 38.145.208.237:8452 | ✓ 1238ms | 否 | ✓ 1307ms | ✓ 775ms | ✓ 1050ms | http |
| 38.145.208.204:8446 | ✓ 1864ms | 否 | ✓ 909ms | ✓ 1944ms | ✓ 1557ms | http |
| 121.230.8.137:1080 | ✓ 1048ms | ✓ 1041ms | ✓ 902ms | ✓ 1826ms | ✓ 1233ms | http |
| 38.145.218.210:8444 | ✓ 1432ms | ✓ 1732ms | 否 | ✓ 861ms | ✓ 780ms | http |
| 38.145.208.192:8453 | ✓ 1810ms | 否 | ✓ 1911ms | ✓ 1266ms | ✓ 1483ms | http |
| 45.136.130.194:8445 | ✓ 520ms | 否 | ✓ 663ms | ✓ 1516ms | 否 | http |
| 139.159.99.242:8080 | ✓ 789ms | ✓ 1040ms | ✓ 1175ms | 否 | 否 | http |
| 120.55.163.237:10086 | ✓ 686ms | ✓ 726ms | ✓ 688ms | ✓ 768ms | ✓ 587ms | http |
| 113.176.92.71:3128 | ✓ 1756ms | 否 | ✓ 1160ms | ✓ 1074ms | ✓ 1333ms | http |

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
