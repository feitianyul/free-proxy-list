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

最后更新：2026-04-18 12:27:10 UTC（2026-04-18 20:27:10 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 1297ms | ✓ 991ms | ✓ 1202ms | ✓ 1232ms | ✓ 761ms | http |
| 1.231.81.166:3128 | ✓ 1428ms | ✓ 931ms | ✓ 637ms | ✓ 904ms | ✓ 654ms | http |
| 149.51.42.10:3128 | ✓ 1345ms | ✓ 1407ms | 否 | ✓ 1875ms | 否 | http |
| 81.30.156.115:8080 | ✓ 636ms | ✓ 1613ms | ✓ 1514ms | ✓ 1744ms | ✓ 1640ms | http |
| 106.10.55.212:1121 | ✓ 1433ms | ✓ 1817ms | ✓ 1142ms | ✓ 1244ms | ✓ 1603ms | http |
| 113.160.132.26:8080 | ✓ 1780ms | ✓ 1629ms | ✓ 838ms | ✓ 1418ms | ✓ 1290ms | http |
| 162.19.253.202:8443 | ✓ 1164ms | 否 | ✓ 1407ms | 否 | ✓ 1979ms | http |
| 160.119.69.7:8080 | ✓ 1981ms | 否 | ✓ 1016ms | 否 | ✓ 1835ms | http |
| 42.101.8.101:8888 | ✓ 1162ms | ✓ 1372ms | ✓ 1301ms | 否 | ✓ 1438ms | http |
| 195.26.224.49:3128 | ✓ 1202ms | 否 | ✓ 1537ms | 否 | ✓ 1829ms | http |
| 149.51.42.10:8080 | ✓ 1011ms | ✓ 1438ms | 否 | ✓ 1467ms | 否 | http |
| 133.18.123.225:26021 | ✓ 507ms | 否 | ✓ 781ms | ✓ 815ms | ✓ 836ms | http |
| 188.246.224.49:7890 | ✓ 784ms | 否 | ✓ 884ms | 否 | ✓ 1902ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 833ms | ✓ 1802ms | ✓ 1238ms | http |
| 185.138.116.150:8080 | ✓ 1426ms | 否 | ✓ 1224ms | ✓ 1794ms | ✓ 1476ms | http |
| 177.93.132.244:3128 | ✓ 969ms | 否 | ✓ 839ms | 否 | ✓ 1758ms | http |
| 84.47.150.125:1080 | ✓ 1188ms | 否 | ✓ 1852ms | 否 | ✓ 1759ms | http |
| 159.89.191.221:3128 | 否 | 否 | ✓ 559ms | ✓ 1504ms | ✓ 1276ms | http |
| 14.247.76.52:8080 | ✓ 1486ms | 否 | 否 | ✓ 1159ms | ✓ 1172ms | http |
| 161.97.184.191:8080 | 否 | ✓ 1995ms | ✓ 1511ms | 否 | ✓ 1545ms | http |
| 59.46.216.131:30001 | ✓ 1052ms | 否 | ✓ 1146ms | ✓ 1279ms | ✓ 1024ms | http |
| 218.108.131.186:17890 | ✓ 791ms | ✓ 1002ms | ✓ 821ms | ✓ 1064ms | ✓ 875ms | http |
| 45.12.151.226:2829 | ✓ 765ms | 否 | ✓ 1396ms | 否 | ✓ 1535ms | http |
| 8.219.195.129:1080 | ✓ 1641ms | 否 | ✓ 982ms | ✓ 1004ms | ✓ 816ms | http |
| 117.122.240.82:3338 | ✓ 1197ms | 否 | ✓ 1022ms | 否 | ✓ 1459ms | http |
| 84.47.150.126:1080 | ✓ 1417ms | 否 | ✓ 1886ms | 否 | ✓ 1847ms | http |
| 168.144.75.9:3128 | 否 | 否 | ✓ 1582ms | ✓ 1907ms | ✓ 1635ms | http |
| 116.58.161.203:26021 | ✓ 1467ms | 否 | ✓ 1847ms | ✓ 1113ms | 否 | http |
| 213.32.85.26:3128 | ✓ 681ms | 否 | ✓ 692ms | ✓ 1951ms | ✓ 1500ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1627ms | ✓ 1136ms | ✓ 928ms | http |
| 43.167.192.85:8080 | ✓ 1523ms | 否 | ✓ 963ms | ✓ 815ms | ✓ 682ms | http |
| 116.171.106.15:3443 | 否 | ✓ 1610ms | ✓ 1507ms | ✓ 1865ms | 否 | http |
| 89.111.174.221:8080 | ✓ 986ms | 否 | ✓ 1644ms | 否 | ✓ 1820ms | http |
| 45.76.207.177:40000 | ✓ 1300ms | 否 | ✓ 1586ms | ✓ 1504ms | ✓ 1250ms | http |
| 91.99.15.45:2095 | ✓ 766ms | ✓ 1788ms | 否 | 否 | ✓ 1720ms | http |
| 212.58.132.5:8888 | ✓ 1764ms | 否 | ✓ 1263ms | ✓ 1777ms | ✓ 1610ms | http |
| 117.236.124.166:3128 | ✓ 1822ms | 否 | ✓ 1933ms | 否 | ✓ 1748ms | http |
| 45.140.147.155:1082 | ✓ 868ms | 否 | 否 | ✓ 1641ms | ✓ 1375ms | http |
| 38.59.240.157:12345 | ✓ 643ms | ✓ 1260ms | 否 | ✓ 1006ms | ✓ 770ms | http |
| 210.223.44.230:3128 | ✓ 814ms | ✓ 1471ms | ✓ 1482ms | ✓ 1572ms | ✓ 1623ms | http |
| 150.107.140.238:3128 | ✓ 1644ms | 否 | ✓ 837ms | 否 | ✓ 840ms | http |
| 43.132.188.134:443 | 否 | 否 | ✓ 1371ms | ✓ 1883ms | ✓ 1824ms | http |
| 120.92.212.16:8890 | ✓ 1142ms | ✓ 1111ms | 否 | ✓ 1908ms | ✓ 1206ms | http |
| 120.92.212.16:7890 | ✓ 1211ms | ✓ 1223ms | ✓ 1914ms | ✓ 1427ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1284ms | ✓ 1258ms | ✓ 1109ms | ✓ 1301ms | ✓ 1585ms | http |
| 194.104.9.38:3128 | 否 | ✓ 1735ms | ✓ 1522ms | ✓ 1848ms | ✓ 1322ms | http |
| 103.53.77.179:8050 | ✓ 1716ms | 否 | 否 | ✓ 1398ms | ✓ 1540ms | http |
| 45.67.202.178:1080 | ✓ 860ms | 否 | ✓ 1883ms | 否 | ✓ 1897ms | http |
| 91.233.223.147:3128 | ✓ 1064ms | 否 | ✓ 1184ms | 否 | ✓ 1902ms | http |
| 70.61.188.34:3128 | ✓ 1076ms | 否 | ✓ 1539ms | ✓ 1570ms | ✓ 1030ms | http |
| 120.92.108.86:7890 | ✓ 1963ms | 否 | ✓ 1665ms | 否 | ✓ 1962ms | http |
| 101.32.243.189:80 | ✓ 1079ms | 否 | ✓ 1673ms | ✓ 1339ms | ✓ 1513ms | http |
| 91.193.240.157:9877 | ✓ 1325ms | 否 | ✓ 1340ms | 否 | ✓ 1877ms | http |
| 60.249.94.208:3128 | 否 | ✓ 1115ms | ✓ 738ms | ✓ 954ms | ✓ 945ms | http |
| 138.124.99.216:8888 | ✓ 1090ms | 否 | ✓ 1459ms | 否 | ✓ 1841ms | http |
| 185.40.7.206:3128 | 否 | ✓ 1638ms | ✓ 1600ms | 否 | ✓ 1838ms | http |
| 171.248.24.178:8080 | ✓ 1178ms | 否 | ✓ 1251ms | ✓ 1568ms | ✓ 1242ms | http |
| 124.156.179.148:3128 | ✓ 656ms | ✓ 1900ms | ✓ 643ms | ✓ 765ms | ✓ 617ms | http |
| 122.117.203.252:3128 | ✓ 711ms | ✓ 1204ms | ✓ 729ms | ✓ 898ms | ✓ 670ms | http |
| 61.171.66.158:3128 | ✓ 863ms | ✓ 1589ms | ✓ 929ms | 否 | ✓ 944ms | http |
| 175.145.210.143:8080 | ✓ 1660ms | 否 | 否 | ✓ 1332ms | ✓ 1509ms | http |
| 47.74.226.8:5001 | ✓ 1382ms | 否 | ✓ 984ms | 否 | ✓ 1260ms | http |
| 5.129.233.140:3128 | ✓ 1822ms | 否 | ✓ 1531ms | ✓ 1479ms | ✓ 1156ms | http |
| 133.18.110.87:1081 | 否 | ✓ 1187ms | ✓ 1698ms | 否 | ✓ 676ms | http |
| 94.131.118.129:1081 | ✓ 1106ms | ✓ 1649ms | ✓ 1336ms | 否 | 否 | http |
| 218.153.163.186:8451 | 否 | 否 | ✓ 1417ms | ✓ 1636ms | ✓ 1319ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1328ms | ✓ 1118ms | ✓ 1498ms | ✓ 1262ms | http |
| 121.230.8.136:1080 | 否 | ✓ 1486ms | ✓ 1351ms | 否 | ✓ 1364ms | http |
| 45.186.6.104:3128 | ✓ 1490ms | ✓ 1942ms | ✓ 1662ms | 否 | 否 | http |
| 103.163.80.25:8081 | ✓ 1961ms | 否 | 否 | ✓ 1578ms | ✓ 1213ms | http |
| 157.66.57.98:8080 | ✓ 1832ms | 否 | 否 | ✓ 1931ms | ✓ 1746ms | http |
| 103.184.99.194:8080 | ✓ 1759ms | 否 | ✓ 1703ms | ✓ 1359ms | ✓ 1324ms | http |
| 124.121.2.162:8080 | ✓ 1742ms | 否 | ✓ 1731ms | ✓ 1515ms | ✓ 1463ms | http |
| 157.230.178.216:8088 | ✓ 1011ms | 否 | 否 | ✓ 1975ms | ✓ 1967ms | http |
| 108.181.201.118:1234 | ✓ 1297ms | ✓ 1177ms | ✓ 887ms | ✓ 1236ms | ✓ 1140ms | http |
| 46.101.95.183:8888 | ✓ 1276ms | 否 | ✓ 1357ms | ✓ 1745ms | ✓ 1369ms | http |
| 82.148.18.242:443 | ✓ 1364ms | 否 | ✓ 1539ms | 否 | ✓ 1866ms | http |
| 104.194.95.90:8080 | ✓ 1502ms | 否 | 否 | ✓ 1285ms | ✓ 980ms | http |
| 103.113.70.189:1082 | ✓ 391ms | 否 | ✓ 1298ms | ✓ 1477ms | 否 | http |
| 103.113.70.189:1081 | ✓ 388ms | 否 | ✓ 1295ms | ✓ 1441ms | 否 | http |
| 12.89.176.82:3128 | ✓ 1106ms | 否 | ✓ 1213ms | ✓ 1839ms | ✓ 1045ms | http |
| 61.52.131.172:8443 | ✓ 811ms | ✓ 1088ms | ✓ 914ms | ✓ 1220ms | ✓ 944ms | http |
| 139.227.17.70:17890 | ✓ 1022ms | ✓ 1001ms | ✓ 1174ms | ✓ 1062ms | ✓ 1134ms | http |
| 103.135.102.161:8081 | 否 | 否 | ✓ 1849ms | ✓ 1484ms | ✓ 901ms | http |
| 160.250.5.22:1 | ✓ 1556ms | 否 | ✓ 1680ms | ✓ 1196ms | ✓ 966ms | http |

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
