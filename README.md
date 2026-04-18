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

最后更新：2026-04-18 09:46:36 UTC（2026-04-18 17:46:36 UTC+8）

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
| 1.231.81.166:3128 | ✓ 1719ms | ✓ 1219ms | ✓ 955ms | ✓ 1193ms | ✓ 936ms | http |
| 218.108.131.186:17890 | ✓ 1022ms | ✓ 1240ms | ✓ 1062ms | ✓ 1318ms | ✓ 1069ms | http |
| 34.96.238.40:8080 | ✓ 1407ms | ✓ 1357ms | 否 | ✓ 1181ms | ✓ 1307ms | http |
| 91.99.15.45:2095 | ✓ 457ms | 否 | 否 | ✓ 1811ms | ✓ 1582ms | http |
| 185.138.116.150:8080 | ✓ 1625ms | ✓ 1687ms | ✓ 1580ms | ✓ 1650ms | ✓ 1415ms | http |
| 157.230.178.216:8088 | 否 | 否 | ✓ 1316ms | ✓ 1649ms | ✓ 1457ms | http |
| 103.85.113.66:9999 | ✓ 1498ms | ✓ 1342ms | 否 | ✓ 1833ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1492ms | ✓ 1326ms | ✓ 1792ms | ✓ 1102ms | http |
| 116.58.161.203:26021 | ✓ 1696ms | 否 | 否 | ✓ 1801ms | ✓ 1243ms | http |
| 45.76.207.177:40000 | ✓ 1696ms | 否 | 否 | ✓ 1975ms | ✓ 1288ms | http |
| 149.51.42.10:8080 | ✓ 261ms | ✓ 974ms | 否 | ✓ 958ms | 否 | http |
| 149.51.42.10:3128 | ✓ 200ms | ✓ 1074ms | 否 | ✓ 961ms | 否 | http |
| 5.104.87.17:8050 | ✓ 702ms | 否 | ✓ 737ms | ✓ 1142ms | ✓ 904ms | http |
| 188.246.224.49:7890 | ✓ 585ms | 否 | ✓ 482ms | 否 | ✓ 1023ms | http |
| 133.18.123.225:26021 | ✓ 724ms | ✓ 1659ms | ✓ 1290ms | ✓ 1064ms | ✓ 903ms | http |
| 159.89.191.221:3128 | ✓ 258ms | ✓ 1285ms | ✓ 1163ms | 否 | 否 | http |
| 213.32.85.26:3128 | ✓ 926ms | 否 | ✓ 1896ms | 否 | ✓ 1625ms | http |
| 212.58.132.5:8888 | ✓ 1540ms | 否 | ✓ 1431ms | ✓ 1460ms | ✓ 1141ms | http |
| 152.32.132.190:7890 | ✓ 1563ms | 否 | ✓ 1004ms | 否 | ✓ 1627ms | http |
| 168.144.75.9:3128 | ✓ 1609ms | 否 | ✓ 1838ms | ✓ 1881ms | ✓ 1211ms | http |
| 91.233.223.147:3128 | ✓ 762ms | 否 | ✓ 723ms | ✓ 1833ms | ✓ 1455ms | http |
| 177.93.132.244:3128 | ✓ 639ms | 否 | ✓ 644ms | 否 | ✓ 1606ms | http |
| 208.87.243.199:7878 | ✓ 718ms | ✓ 1708ms | 否 | ✓ 1530ms | 否 | http |
| 35.225.22.61:80 | ✓ 930ms | 否 | ✓ 1343ms | ✓ 1028ms | ✓ 939ms | http |
| 162.19.253.202:8443 | ✓ 1447ms | ✓ 1731ms | ✓ 920ms | 否 | ✓ 1542ms | http |
| 45.12.151.226:2829 | 否 | ✓ 1512ms | ✓ 983ms | 否 | ✓ 1657ms | http |
| 120.92.108.86:7890 | ✓ 1468ms | 否 | 否 | ✓ 1970ms | ✓ 1501ms | http |
| 120.92.212.16:8890 | ✓ 1977ms | 否 | 否 | ✓ 1512ms | ✓ 1638ms | http |
| 38.59.240.157:12345 | ✓ 1507ms | 否 | 否 | ✓ 1390ms | ✓ 835ms | http |
| 201.144.20.238:3128 | ✓ 1215ms | 否 | ✓ 1910ms | 否 | ✓ 1153ms | http |
| 84.47.150.125:1080 | ✓ 935ms | 否 | ✓ 1964ms | 否 | ✓ 1701ms | http |
| 178.238.117.178:8080 | ✓ 962ms | 否 | ✓ 1537ms | 否 | ✓ 1813ms | http |
| 164.92.166.127:3128 | ✓ 1451ms | ✓ 1629ms | ✓ 723ms | ✓ 1231ms | ✓ 912ms | http |
| 94.158.219.111:3128 | 否 | 否 | ✓ 1026ms | ✓ 1712ms | ✓ 1478ms | http |
| 59.46.216.131:30001 | ✓ 1227ms | ✓ 1623ms | ✓ 1364ms | 否 | 否 | http |
| 103.113.70.189:1082 | ✓ 1509ms | ✓ 1463ms | ✓ 1244ms | 否 | 否 | http |
| 201.71.24.65:8082 | ✓ 1322ms | 否 | ✓ 1683ms | 否 | ✓ 1959ms | http |
| 45.186.6.104:3128 | ✓ 1106ms | ✓ 1634ms | ✓ 1623ms | 否 | 否 | http |
| 164.163.42.5:10000 | ✓ 1161ms | 否 | ✓ 1128ms | 否 | ✓ 1749ms | http |
| 218.60.0.214:80 | ✓ 1141ms | 否 | ✓ 1221ms | ✓ 1469ms | 否 | http |
| 82.148.18.242:443 | ✓ 896ms | 否 | ✓ 1378ms | ✓ 1916ms | 否 | http |
| 144.31.27.49:1080 | ✓ 889ms | 否 | 否 | ✓ 1711ms | ✓ 1431ms | http |
| 138.124.99.216:8888 | ✓ 1078ms | ✓ 1795ms | ✓ 1493ms | 否 | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1747ms | ✓ 911ms | ✓ 1141ms | ✓ 937ms | http |
| 38.180.2.107:3128 | ✓ 1390ms | 否 | ✓ 1727ms | 否 | ✓ 1676ms | http |
| 117.236.124.166:3128 | ✓ 1722ms | 否 | ✓ 1537ms | 否 | ✓ 1749ms | http |
| 223.84.151.86:30005 | ✓ 1570ms | ✓ 1481ms | ✓ 1391ms | ✓ 1756ms | ✓ 1557ms | http |
| 120.92.212.16:7890 | ✓ 1366ms | ✓ 1343ms | ✓ 1153ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 437ms | 否 | ✓ 1219ms | ✓ 1954ms | ✓ 1265ms | http |
| 51.79.207.21:8080 | ✓ 1114ms | 否 | ✓ 1811ms | ✓ 1923ms | ✓ 1706ms | http |
| 91.107.124.215:3128 | ✓ 661ms | ✓ 1803ms | ✓ 670ms | ✓ 1976ms | ✓ 1592ms | http |
| 103.113.70.189:1081 | ✓ 975ms | 否 | ✓ 1210ms | 否 | ✓ 1078ms | http |
| 101.32.243.189:80 | ✓ 1635ms | 否 | ✓ 1813ms | 否 | ✓ 1687ms | http |
| 160.238.65.6:3128 | ✓ 372ms | ✓ 1167ms | ✓ 641ms | ✓ 1180ms | ✓ 1713ms | http |
| 160.238.65.3:3128 | ✓ 404ms | ✓ 1700ms | ✓ 403ms | ✓ 1390ms | ✓ 1174ms | http |
| 160.238.65.5:3128 | ✓ 375ms | ✓ 1935ms | ✓ 401ms | ✓ 1269ms | ✓ 1086ms | http |
| 160.238.65.4:3128 | ✓ 374ms | 否 | ✓ 459ms | ✓ 1218ms | ✓ 1024ms | http |
| 160.238.65.2:3128 | ✓ 374ms | ✓ 1343ms | ✓ 460ms | ✓ 1204ms | ✓ 1693ms | http |
| 160.238.65.7:3128 | ✓ 380ms | 否 | ✓ 665ms | ✓ 1214ms | ✓ 943ms | http |
| 160.238.65.8:3128 | ✓ 390ms | 否 | ✓ 655ms | ✓ 1211ms | ✓ 947ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 1639ms | ✓ 1412ms | ✓ 1415ms | http |
| 45.140.147.155:1082 | ✓ 1199ms | ✓ 1291ms | ✓ 1711ms | ✓ 1678ms | ✓ 1188ms | http |
| 160.238.65.9:3128 | ✓ 964ms | ✓ 1480ms | ✓ 665ms | ✓ 1880ms | ✓ 1023ms | http |
| 51.81.6.158:3128 | ✓ 1538ms | ✓ 1637ms | ✓ 1418ms | 否 | 否 | http |
| 82.114.228.67:1080 | ✓ 630ms | 否 | ✓ 1340ms | ✓ 1832ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1621ms | 否 | ✓ 1689ms | ✓ 1499ms | ✓ 1537ms | http |
| 178.63.155.151:9000 | ✓ 1016ms | ✓ 1733ms | 否 | 否 | ✓ 1479ms | http |
| 144.31.25.69:21064 | ✓ 478ms | 否 | ✓ 907ms | 否 | ✓ 1788ms | http |
| 84.47.150.126:1080 | ✓ 1118ms | ✓ 1949ms | 否 | 否 | ✓ 1856ms | http |
| 45.140.147.155:1081 | ✓ 1083ms | ✓ 1621ms | ✓ 1409ms | ✓ 1191ms | ✓ 831ms | http |
| 83.219.250.8:62920 | ✓ 559ms | ✓ 1307ms | ✓ 1265ms | ✓ 1800ms | ✓ 1728ms | http |
| 196.206.223.88:1221 | ✓ 1616ms | 否 | ✓ 1046ms | 否 | ✓ 1610ms | http |
| 196.206.216.164:1221 | ✓ 1745ms | 否 | ✓ 1009ms | 否 | ✓ 1864ms | http |
| 8.219.195.129:1080 | ✓ 1812ms | 否 | ✓ 935ms | ✓ 1251ms | ✓ 1089ms | http |
| 8.137.112.117:3128 | ✓ 1705ms | 否 | ✓ 1953ms | ✓ 1465ms | ✓ 1213ms | http |
| 103.3.58.162:8088 | 否 | 否 | ✓ 1717ms | ✓ 1792ms | ✓ 1771ms | http |
| 204.48.31.203:80 | ✓ 822ms | ✓ 1728ms | ✓ 792ms | ✓ 926ms | ✓ 1767ms | http |
| 160.119.69.7:8080 | ✓ 1865ms | 否 | ✓ 1837ms | ✓ 1791ms | ✓ 1437ms | http |
| 94.131.118.129:1081 | ✓ 1927ms | 否 | ✓ 1838ms | 否 | ✓ 1849ms | http |
| 62.113.119.14:8080 | ✓ 1918ms | ✓ 1540ms | ✓ 647ms | ✓ 1790ms | ✓ 1172ms | http |
| 106.10.55.212:1121 | ✓ 1547ms | ✓ 1310ms | ✓ 1459ms | ✓ 1774ms | ✓ 1396ms | http |
| 61.52.131.172:8443 | ✓ 1856ms | ✓ 1459ms | 否 | ✓ 1440ms | ✓ 1171ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 961ms | ✓ 1346ms | ✓ 1102ms | http |
| 43.132.188.134:443 | ✓ 1590ms | 否 | 否 | ✓ 1875ms | ✓ 1645ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1670ms | ✓ 1705ms | ✓ 1123ms | http |

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
