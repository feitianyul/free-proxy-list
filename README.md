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

最后更新：2026-04-21 16:10:15 UTC（2026-04-22 00:10:15 UTC+8）

**代理总数：94**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 94 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1268ms | 否 | ✓ 1556ms | ✓ 866ms | ✓ 679ms | http |
| 152.42.208.139:8118 | ✓ 1434ms | 否 | ✓ 1197ms | ✓ 1199ms | ✓ 784ms | http |
| 113.160.132.26:8080 | ✓ 1432ms | ✓ 1324ms | 否 | ✓ 1254ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1434ms | 否 | ✓ 1509ms | ✓ 1632ms | ✓ 1297ms | http |
| 46.101.95.183:8888 | ✓ 1646ms | ✓ 1874ms | ✓ 1743ms | 否 | ✓ 1578ms | http |
| 81.30.156.115:8080 | ✓ 1645ms | ✓ 1987ms | 否 | 否 | ✓ 1796ms | http |
| 35.225.22.61:80 | 否 | ✓ 1367ms | ✓ 446ms | ✓ 1100ms | ✓ 1028ms | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 825ms | ✓ 1761ms | ✓ 989ms | http |
| 162.19.253.202:8443 | ✓ 1183ms | 否 | ✓ 1752ms | 否 | ✓ 1853ms | http |
| 223.84.151.86:30005 | ✓ 1452ms | ✓ 1398ms | ✓ 1296ms | 否 | ✓ 1535ms | http |
| 106.10.55.212:1121 | ✓ 953ms | 否 | 否 | ✓ 1493ms | ✓ 1949ms | http |
| 188.246.224.49:7890 | 否 | 否 | ✓ 1861ms | ✓ 1653ms | ✓ 1642ms | http |
| 208.87.243.199:7878 | ✓ 1381ms | 否 | ✓ 1261ms | 否 | ✓ 1347ms | http |
| 128.199.116.219:9090 | ✓ 1406ms | 否 | ✓ 1400ms | ✓ 1405ms | 否 | http |
| 128.199.113.85:9090 | ✓ 1397ms | 否 | ✓ 907ms | ✓ 1137ms | ✓ 1398ms | http |
| 89.208.106.138:10808 | ✓ 1286ms | 否 | 否 | ✓ 1930ms | ✓ 1380ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1339ms | ✓ 1941ms | ✓ 1722ms | 否 | http |
| 91.99.15.45:2095 | ✓ 1305ms | 否 | ✓ 1801ms | ✓ 1896ms | ✓ 1724ms | http |
| 185.138.116.150:8080 | ✓ 1432ms | 否 | ✓ 1251ms | 否 | ✓ 1318ms | http |
| 91.233.223.147:3128 | ✓ 1330ms | 否 | ✓ 1070ms | 否 | ✓ 1772ms | http |
| 78.11.96.22:8888 | ✓ 1302ms | ✓ 1866ms | ✓ 1974ms | ✓ 1779ms | ✓ 1906ms | http |
| 8.219.195.129:1080 | ✓ 682ms | 否 | ✓ 723ms | ✓ 1014ms | ✓ 793ms | http |
| 34.71.229.255:3128 | ✓ 822ms | 否 | 否 | ✓ 1553ms | ✓ 1302ms | http |
| 42.101.8.101:8888 | 否 | ✓ 1515ms | ✓ 1315ms | ✓ 1743ms | ✓ 1599ms | http |
| 115.231.181.40:8128 | ✓ 1161ms | 否 | ✓ 961ms | ✓ 1481ms | ✓ 998ms | http |
| 34.96.238.40:8080 | ✓ 870ms | ✓ 1208ms | 否 | ✓ 903ms | ✓ 1265ms | http |
| 84.47.150.125:1080 | ✓ 1272ms | 否 | ✓ 1345ms | 否 | ✓ 1890ms | http |
| 14.143.222.113:57788 | ✓ 1746ms | 否 | ✓ 1119ms | ✓ 1566ms | 否 | http |
| 168.144.75.9:3128 | ✓ 1755ms | 否 | ✓ 1929ms | 否 | ✓ 1429ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1291ms | 否 | ✓ 1136ms | ✓ 984ms | http |
| 45.140.147.82:1081 | ✓ 710ms | ✓ 1554ms | ✓ 1618ms | ✓ 1782ms | ✓ 1220ms | http |
| 45.140.147.82:1082 | ✓ 715ms | ✓ 1534ms | ✓ 1631ms | ✓ 1741ms | ✓ 1296ms | http |
| 59.46.216.131:30001 | ✓ 1711ms | 否 | ✓ 1639ms | ✓ 1364ms | 否 | http |
| 20.78.26.206:8561 | ✓ 442ms | ✓ 833ms | ✓ 753ms | ✓ 755ms | ✓ 608ms | http |
| 148.153.56.51:80 | ✓ 549ms | ✓ 824ms | 否 | ✓ 1008ms | ✓ 558ms | http |
| 20.210.39.153:8561 | ✓ 441ms | ✓ 1769ms | ✓ 431ms | ✓ 748ms | ✓ 593ms | http |
| 20.78.118.91:8561 | ✓ 441ms | ✓ 1845ms | ✓ 469ms | ✓ 727ms | ✓ 580ms | http |
| 159.203.220.84:3128 | ✓ 832ms | 否 | 否 | ✓ 966ms | ✓ 687ms | http |
| 118.193.44.243:3128 | ✓ 704ms | 否 | ✓ 1027ms | ✓ 830ms | ✓ 646ms | http |
| 210.77.29.244:6478 | ✓ 786ms | ✓ 1944ms | ✓ 896ms | ✓ 1104ms | ✓ 799ms | http |
| 121.230.9.160:1080 | ✓ 1717ms | 否 | ✓ 1899ms | ✓ 1258ms | ✓ 1769ms | http |
| 103.155.199.15:8080 | ✓ 1271ms | 否 | ✓ 1216ms | 否 | ✓ 1395ms | http |
| 103.46.11.156:3125 | ✓ 1290ms | 否 | ✓ 1402ms | ✓ 1454ms | 否 | http |
| 103.97.224.219:8181 | ✓ 1353ms | 否 | 否 | ✓ 1708ms | ✓ 1782ms | http |
| 70.61.188.34:3128 | ✓ 1388ms | 否 | ✓ 1857ms | ✓ 1522ms | ✓ 1364ms | http |
| 62.113.119.14:8080 | ✓ 1059ms | ✓ 1750ms | ✓ 952ms | ✓ 1771ms | ✓ 1275ms | http |
| 149.51.42.10:3128 | ✓ 731ms | ✓ 1871ms | 否 | ✓ 1738ms | 否 | http |
| 45.76.207.177:40000 | ✓ 706ms | 否 | ✓ 643ms | ✓ 1006ms | ✓ 785ms | http |
| 139.159.97.82:10900 | ✓ 1050ms | 否 | ✓ 979ms | ✓ 1382ms | ✓ 1098ms | http |
| 159.89.191.221:3128 | ✓ 489ms | 否 | ✓ 596ms | ✓ 1616ms | ✓ 1447ms | http |
| 187.189.63.149:8080 | ✓ 832ms | 否 | ✓ 700ms | ✓ 1605ms | 否 | http |
| 149.51.42.10:8080 | ✓ 1321ms | ✓ 1892ms | 否 | ✓ 1755ms | 否 | http |
| 160.238.65.6:3128 | ✓ 1075ms | ✓ 1670ms | 否 | 否 | ✓ 1914ms | http |
| 47.74.226.8:5001 | ✓ 1419ms | ✓ 1369ms | ✓ 886ms | ✓ 1218ms | 否 | http |
| 160.238.65.4:3128 | ✓ 698ms | 否 | ✓ 979ms | 否 | ✓ 1853ms | http |
| 160.238.65.9:3128 | ✓ 620ms | ✓ 1594ms | ✓ 1161ms | 否 | ✓ 1861ms | http |
| 160.238.65.7:3128 | ✓ 685ms | ✓ 1449ms | ✓ 1551ms | 否 | ✓ 1961ms | http |
| 160.238.65.3:3128 | ✓ 632ms | 否 | ✓ 641ms | ✓ 1669ms | ✓ 1181ms | http |
| 130.61.139.145:3128 | ✓ 1133ms | 否 | 否 | ✓ 1935ms | ✓ 1618ms | http |
| 160.238.65.2:3128 | ✓ 692ms | ✓ 1537ms | ✓ 919ms | ✓ 1563ms | ✓ 1290ms | http |
| 103.113.70.189:1081 | ✓ 1298ms | ✓ 1168ms | ✓ 401ms | 否 | 否 | http |
| 43.132.188.134:443 | 否 | ✓ 871ms | 否 | ✓ 776ms | ✓ 1443ms | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 1727ms | ✓ 1724ms | ✓ 1146ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 1735ms | ✓ 1775ms | ✓ 1210ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 1738ms | ✓ 1769ms | ✓ 1218ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 1739ms | ✓ 1768ms | ✓ 1210ms | http |
| 45.129.141.174:3128 | ✓ 1293ms | 否 | ✓ 1858ms | ✓ 1977ms | ✓ 1778ms | http |
| 120.92.212.16:8890 | ✓ 1601ms | ✓ 1606ms | 否 | ✓ 1699ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1686ms | 否 | ✓ 730ms | 否 | ✓ 1908ms | http |
| 121.232.73.221:1080 | ✓ 1088ms | ✓ 1861ms | ✓ 1279ms | 否 | ✓ 949ms | http |
| 120.92.212.16:7890 | ✓ 1193ms | ✓ 1217ms | ✓ 1200ms | 否 | 否 | http |
| 84.47.150.126:1080 | ✓ 1109ms | 否 | ✓ 1575ms | 否 | ✓ 1888ms | http |
| 120.92.108.86:7890 | ✓ 1198ms | 否 | ✓ 1843ms | ✓ 1884ms | 否 | http |
| 210.223.44.230:3128 | ✓ 756ms | 否 | ✓ 1328ms | ✓ 1084ms | ✓ 886ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 1837ms | ✓ 939ms | ✓ 1779ms | http |
| 49.0.26.215:8080 | ✓ 1822ms | 否 | ✓ 1241ms | ✓ 1875ms | 否 | http |
| 104.248.151.93:9090 | ✓ 1412ms | 否 | ✓ 1927ms | 否 | ✓ 1642ms | http |
| 36.141.21.200:7890 | 否 | ✓ 1678ms | ✓ 904ms | 否 | ✓ 954ms | http |
| 45.12.151.226:2829 | ✓ 1233ms | ✓ 1702ms | ✓ 1243ms | 否 | 否 | http |
| 20.210.76.178:8561 | ✓ 1243ms | ✓ 1600ms | ✓ 1233ms | ✓ 1198ms | ✓ 1501ms | http |
| 20.18.193.135:8561 | ✓ 1252ms | ✓ 1864ms | ✓ 1123ms | ✓ 1158ms | ✓ 1388ms | http |
| 20.210.76.104:8561 | ✓ 1251ms | ✓ 1613ms | ✓ 1222ms | ✓ 1216ms | ✓ 1482ms | http |
| 20.210.76.175:8561 | ✓ 1255ms | 否 | ✓ 1036ms | ✓ 1434ms | ✓ 1208ms | http |
| 20.27.15.49:8561 | ✓ 1261ms | 否 | ✓ 1041ms | ✓ 1443ms | ✓ 1214ms | http |
| 116.171.106.26:3443 | 否 | ✓ 1451ms | ✓ 1356ms | ✓ 1678ms | ✓ 1753ms | http |
| 150.249.255.91:3128 | ✓ 484ms | ✓ 1371ms | ✓ 683ms | ✓ 1959ms | ✓ 775ms | http |
| 42.200.76.16:3888 | ✓ 609ms | 否 | ✓ 607ms | ✓ 817ms | ✓ 649ms | http |
| 152.70.91.193:40000 | ✓ 1396ms | 否 | ✓ 1630ms | 否 | ✓ 1202ms | http |
| 117.122.240.82:3338 | 否 | ✓ 1922ms | ✓ 1683ms | ✓ 1561ms | ✓ 880ms | http |
| 112.163.160.93:3128 | 否 | 否 | ✓ 1682ms | ✓ 963ms | ✓ 737ms | http |
| 176.124.220.172:3128 | ✓ 1081ms | 否 | ✓ 1696ms | 否 | ✓ 1718ms | http |
| 61.52.131.172:8443 | ✓ 803ms | ✓ 1127ms | ✓ 930ms | ✓ 1220ms | ✓ 921ms | http |
| 157.20.128.247:3125 | ✓ 1856ms | 否 | ✓ 1492ms | ✓ 1312ms | 否 | http |
| 8.219.97.248:80 | ✓ 1328ms | 否 | ✓ 1111ms | ✓ 1632ms | 否 | http |

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
