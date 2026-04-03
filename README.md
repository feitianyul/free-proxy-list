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

最后更新：2026-04-03 23:36:50 UTC（2026-04-04 07:36:50 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 353ms | ✓ 1443ms | ✓ 953ms | ✓ 1120ms | 否 | http |
| 218.108.131.186:17890 | ✓ 769ms | ✓ 983ms | ✓ 774ms | ✓ 1057ms | ✓ 845ms | http |
| 147.161.210.140:8800 | ✓ 698ms | 否 | ✓ 1042ms | ✓ 995ms | ✓ 946ms | http |
| 1.231.81.166:3128 | ✓ 751ms | ✓ 1081ms | ✓ 1725ms | ✓ 1298ms | ✓ 1056ms | http |
| 43.167.237.94:3128 | ✓ 459ms | 否 | ✓ 1972ms | ✓ 1282ms | ✓ 581ms | http |
| 111.227.254.9:22222 | ✓ 1024ms | ✓ 1222ms | ✓ 913ms | ✓ 1325ms | ✓ 1064ms | http |
| 111.227.254.12:22222 | ✓ 965ms | ✓ 1211ms | ✓ 1049ms | ✓ 1323ms | ✓ 1106ms | http |
| 113.160.132.26:8080 | ✓ 1434ms | ✓ 1321ms | 否 | ✓ 1205ms | ✓ 921ms | http |
| 167.103.115.102:8800 | ✓ 1100ms | 否 | ✓ 941ms | ✓ 1119ms | ✓ 1186ms | http |
| 160.250.134.143:3128 | ✓ 960ms | 否 | ✓ 1319ms | ✓ 1155ms | ✓ 969ms | http |
| 95.213.217.168:52004 | ✓ 869ms | ✓ 1815ms | 否 | 否 | ✓ 1331ms | http |
| 34.101.184.164:3128 | ✓ 1455ms | 否 | ✓ 1837ms | ✓ 1280ms | ✓ 990ms | http |
| 167.103.34.108:8800 | ✓ 1300ms | ✓ 1936ms | ✓ 1377ms | 否 | ✓ 1676ms | http |
| 167.103.144.127:8800 | ✓ 1908ms | 否 | ✓ 1273ms | ✓ 1471ms | ✓ 1941ms | http |
| 208.87.243.199:7878 | ✓ 397ms | ✓ 972ms | ✓ 584ms | ✓ 953ms | ✓ 1773ms | http |
| 49.229.100.235:8080 | ✓ 1556ms | ✓ 1944ms | ✓ 1646ms | 否 | 否 | http |
| 72.11.151.159:6005 | ✓ 610ms | ✓ 1384ms | ✓ 942ms | ✓ 1315ms | ✓ 922ms | http |
| 174.140.109.250:3128 | 否 | ✓ 1415ms | ✓ 966ms | ✓ 1384ms | ✓ 985ms | http |
| 167.103.31.122:8800 | ✓ 1310ms | 否 | ✓ 1293ms | 否 | ✓ 1529ms | http |
| 120.92.212.16:8890 | ✓ 996ms | 否 | ✓ 1035ms | ✓ 1155ms | ✓ 1178ms | http |
| 120.92.212.16:7890 | ✓ 1422ms | 否 | ✓ 1299ms | 否 | ✓ 1594ms | http |
| 45.136.130.194:8451 | ✓ 1480ms | ✓ 791ms | ✓ 1229ms | 否 | 否 | http |
| 114.237.77.202:1080 | 否 | 否 | ✓ 960ms | ✓ 1186ms | ✓ 892ms | http |
| 101.43.127.100:8877 | ✓ 898ms | ✓ 1550ms | ✓ 763ms | ✓ 1100ms | ✓ 892ms | http |
| 147.161.239.240:8800 | ✓ 1124ms | 否 | ✓ 1031ms | ✓ 1808ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1585ms | 否 | ✓ 890ms | ✓ 1081ms | ✓ 977ms | http |
| 59.46.216.131:30001 | ✓ 899ms | ✓ 1327ms | ✓ 1057ms | ✓ 1247ms | ✓ 1049ms | http |
| 177.234.217.88:999 | ✓ 1356ms | ✓ 1767ms | ✓ 1809ms | ✓ 1990ms | 否 | http |
| 123.57.2.231:2020 | ✓ 1077ms | ✓ 1222ms | ✓ 1061ms | ✓ 1169ms | ✓ 975ms | http |
| 1.225.116.115:1080 | ✓ 1301ms | ✓ 1501ms | ✓ 1355ms | ✓ 1961ms | ✓ 899ms | http |
| 38.145.218.9:8445 | ✓ 1070ms | 否 | ✓ 709ms | ✓ 719ms | ✓ 843ms | http |
| 195.123.209.48:3128 | ✓ 1327ms | ✓ 1750ms | ✓ 1433ms | 否 | ✓ 1954ms | http |
| 217.76.245.80:999 | ✓ 1149ms | ✓ 1662ms | ✓ 1329ms | ✓ 1856ms | ✓ 1576ms | http |
| 159.223.71.162:8080 | ✓ 1453ms | 否 | 否 | ✓ 991ms | ✓ 787ms | http |
| 38.145.220.11:8447 | ✓ 1018ms | ✓ 1692ms | ✓ 198ms | ✓ 847ms | 否 | http |
| 38.34.179.18:8451 | 否 | ✓ 1094ms | ✓ 1824ms | 否 | ✓ 528ms | http |
| 209.126.84.232:8888 | ✓ 821ms | 否 | ✓ 798ms | 否 | ✓ 1568ms | http |
| 103.72.89.22:8097 | ✓ 1603ms | 否 | ✓ 1756ms | ✓ 1944ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1347ms | 否 | ✓ 807ms | ✓ 1880ms | 否 | http |
| 138.197.68.35:4857 | ✓ 492ms | ✓ 1770ms | 否 | ✓ 1435ms | 否 | http |
| 45.136.131.29:8452 | ✓ 594ms | ✓ 1095ms | ✓ 1803ms | ✓ 1353ms | ✓ 669ms | http |
| 46.39.105.157:8080 | ✓ 1328ms | 否 | ✓ 1925ms | 否 | ✓ 1594ms | http |
| 5.104.87.17:8050 | ✓ 1686ms | ✓ 1732ms | ✓ 1194ms | ✓ 865ms | ✓ 702ms | http |
| 27.254.99.183:8118 | 否 | ✓ 1998ms | 否 | ✓ 1242ms | ✓ 1000ms | http |
| 112.163.160.93:3128 | 否 | ✓ 776ms | ✓ 826ms | ✓ 1592ms | 否 | http |
| 147.45.186.28:3128 | ✓ 1337ms | 否 | ✓ 1344ms | ✓ 1844ms | ✓ 1514ms | http |
| 116.80.63.46:7777 | ✓ 1643ms | 否 | ✓ 1988ms | ✓ 1962ms | 否 | http |
| 159.223.71.162:443 | ✓ 690ms | 否 | ✓ 696ms | ✓ 1022ms | ✓ 788ms | http |
| 116.80.96.101:3172 | ✓ 1513ms | 否 | ✓ 1497ms | ✓ 1802ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1162ms | 否 | ✓ 1264ms | ✓ 1592ms | ✓ 1613ms | http |
| 180.103.19.53:1080 | ✓ 1122ms | ✓ 1779ms | ✓ 1059ms | ✓ 1411ms | ✓ 1170ms | http |
| 106.117.208.101:7890 | ✓ 925ms | ✓ 1313ms | ✓ 1119ms | ✓ 1241ms | ✓ 1007ms | http |
| 121.43.189.36:28888 | ✓ 1488ms | ✓ 1878ms | ✓ 1472ms | ✓ 1502ms | ✓ 1427ms | http |
| 101.32.244.83:8080 | ✓ 968ms | ✓ 1467ms | ✓ 896ms | ✓ 1110ms | ✓ 1161ms | http |
| 121.43.196.210:8222 | ✓ 966ms | ✓ 990ms | ✓ 845ms | ✓ 1048ms | ✓ 885ms | http |
| 121.43.196.213:8222 | ✓ 893ms | ✓ 1040ms | ✓ 824ms | ✓ 1056ms | ✓ 870ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1367ms | 否 | ✓ 1904ms | ✓ 1703ms | http |
| 114.55.226.123:10086 | ✓ 1051ms | ✓ 1343ms | ✓ 1012ms | ✓ 1287ms | ✓ 1050ms | http |
| 150.241.71.15:1080 | ✓ 1242ms | ✓ 1942ms | ✓ 1156ms | 否 | 否 | http |
| 38.34.179.174:8453 | ✓ 875ms | ✓ 651ms | ✓ 651ms | 否 | ✓ 1182ms | http |
| 185.191.236.162:3128 | ✓ 1366ms | ✓ 1747ms | 否 | 否 | ✓ 1560ms | http |
| 103.203.233.113:3125 | ✓ 1234ms | 否 | 否 | ✓ 1598ms | ✓ 1391ms | http |
| 38.145.208.215:8444 | ✓ 1075ms | 否 | ✓ 167ms | ✓ 691ms | ✓ 1246ms | http |
| 210.223.44.230:3128 | ✓ 1746ms | 否 | ✓ 641ms | ✓ 1093ms | ✓ 970ms | http |
| 34.96.238.40:8080 | ✓ 1630ms | ✓ 1587ms | 否 | ✓ 1228ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 798ms | ✓ 1557ms | ✓ 1655ms | http |
| 180.130.80.196:9003 | ✓ 1150ms | ✓ 1340ms | ✓ 1294ms | ✓ 1371ms | ✓ 1061ms | http |
| 142.171.95.105:3128 | 否 | 否 | ✓ 506ms | ✓ 1473ms | ✓ 497ms | http |
| 45.136.198.40:3128 | ✓ 1612ms | ✓ 1701ms | ✓ 1747ms | 否 | ✓ 1642ms | http |
| 194.67.99.223:1080 | ✓ 1330ms | ✓ 1753ms | 否 | 否 | ✓ 1586ms | http |
| 101.132.61.121:8888 | ✓ 1213ms | 否 | ✓ 1275ms | ✓ 1331ms | ✓ 1189ms | http |
| 82.114.228.67:1080 | ✓ 1263ms | ✓ 1755ms | 否 | 否 | ✓ 1742ms | http |
| 150.249.255.91:3128 | ✓ 514ms | ✓ 859ms | 否 | 否 | ✓ 1217ms | http |
| 92.119.127.211:6005 | ✓ 1248ms | ✓ 1745ms | ✓ 1816ms | 否 | ✓ 1855ms | http |
| 180.148.25.182:8085 | ✓ 1808ms | 否 | 否 | ✓ 1327ms | ✓ 1349ms | http |
| 125.64.244.100:8889 | ✓ 1581ms | ✓ 1571ms | ✓ 1541ms | ✓ 1893ms | 否 | http |
| 64.227.76.27:1080 | ✓ 1169ms | ✓ 1724ms | 否 | ✓ 1712ms | 否 | http |
| 38.145.220.60:8447 | ✓ 631ms | ✓ 892ms | 否 | ✓ 1524ms | 否 | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1191ms | ✓ 1931ms | ✓ 1231ms | http |
| 217.217.249.160:8080 | ✓ 1411ms | 否 | ✓ 1292ms | 否 | ✓ 1729ms | http |
| 38.34.179.172:8447 | ✓ 825ms | ✓ 618ms | ✓ 193ms | ✓ 683ms | ✓ 1268ms | http |
| 121.230.9.198:1080 | ✓ 1019ms | ✓ 1294ms | ✓ 1289ms | ✓ 1316ms | ✓ 1338ms | http |
| 121.230.8.213:1080 | ✓ 1053ms | ✓ 1420ms | ✓ 1069ms | ✓ 1694ms | ✓ 1254ms | http |
| 172.238.122.148:3128 | ✓ 1559ms | 否 | ✓ 1355ms | ✓ 1988ms | 否 | http |
| 114.237.77.239:1080 | 否 | 否 | ✓ 855ms | ✓ 1224ms | ✓ 993ms | http |
| 203.150.113.117:8080 | ✓ 1474ms | 否 | ✓ 1913ms | ✓ 1746ms | ✓ 1536ms | http |
| 38.34.179.190:8452 | ✓ 275ms | ✓ 637ms | 否 | ✓ 1720ms | ✓ 1040ms | http |
| 183.232.248.73:7890 | 否 | ✓ 1106ms | ✓ 913ms | ✓ 1687ms | 否 | http |
| 121.230.8.208:1080 | 否 | 否 | ✓ 973ms | ✓ 1537ms | ✓ 1162ms | http |
| 38.145.218.216:8449 | 否 | ✓ 1077ms | ✓ 339ms | 否 | ✓ 1503ms | http |
| 3.79.194.222:42707 | ✓ 1274ms | 否 | 否 | ✓ 1882ms | ✓ 1848ms | http |
| 103.113.70.189:1081 | ✓ 716ms | ✓ 1127ms | ✓ 1244ms | ✓ 1152ms | ✓ 901ms | http |

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
