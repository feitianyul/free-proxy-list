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

最后更新：2026-05-03 00:41:44 UTC（2026-05-03 08:41:44 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 97ms | ✓ 957ms | 否 | ✓ 1275ms | ✓ 904ms | http |
| 47.77.216.82:1080 | ✓ 212ms | ✓ 938ms | 否 | ✓ 1184ms | ✓ 1890ms | http |
| 80.92.204.47:1081 | ✓ 880ms | ✓ 1259ms | ✓ 783ms | ✓ 1866ms | ✓ 1256ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1452ms | 否 | ✓ 1762ms | ✓ 1457ms | http |
| 148.230.4.241:999 | ✓ 864ms | ✓ 1985ms | ✓ 698ms | ✓ 1651ms | ✓ 1451ms | http |
| 45.167.124.71:999 | ✓ 1121ms | ✓ 1685ms | ✓ 1316ms | ✓ 1748ms | ✓ 1799ms | http |
| 45.140.147.155:1082 | ✓ 491ms | 否 | ✓ 689ms | 否 | ✓ 1279ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1643ms | ✓ 1986ms | ✓ 1026ms | 否 | http |
| 168.110.52.228:3128 | ✓ 960ms | ✓ 1659ms | 否 | ✓ 1124ms | ✓ 1230ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1542ms | ✓ 1385ms | 否 | ✓ 1167ms | http |
| 160.238.65.7:3128 | ✓ 1481ms | ✓ 1615ms | ✓ 1814ms | 否 | 否 | http |
| 160.238.65.2:3128 | ✓ 1481ms | ✓ 1752ms | ✓ 1681ms | 否 | 否 | http |
| 147.45.178.211:14658 | ✓ 1174ms | ✓ 1702ms | ✓ 813ms | ✓ 1729ms | 否 | http |
| 218.108.131.186:17890 | ✓ 962ms | ✓ 1174ms | ✓ 1000ms | ✓ 1226ms | ✓ 1007ms | http |
| 45.59.122.132:80 | ✓ 756ms | 否 | ✓ 1109ms | 否 | ✓ 1229ms | http |
| 62.60.231.71:56608 | 否 | ✓ 1669ms | ✓ 1432ms | 否 | ✓ 1563ms | http |
| 120.92.108.86:7890 | ✓ 1405ms | 否 | 否 | ✓ 1807ms | ✓ 1479ms | http |
| 45.153.231.229:8080 | ✓ 1647ms | 否 | ✓ 1688ms | 否 | ✓ 1912ms | http |
| 149.51.42.10:3128 | ✓ 1035ms | ✓ 1228ms | 否 | ✓ 1460ms | 否 | http |
| 149.51.42.10:8080 | ✓ 1037ms | ✓ 1260ms | 否 | ✓ 1466ms | 否 | http |
| 109.120.156.122:8090 | ✓ 1251ms | ✓ 1850ms | ✓ 733ms | 否 | ✓ 1850ms | http |
| 212.58.132.5:8888 | ✓ 1292ms | 否 | ✓ 1451ms | ✓ 1572ms | ✓ 1221ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1487ms | 否 | ✓ 1516ms | ✓ 1063ms | http |
| 160.238.65.4:3128 | ✓ 1479ms | ✓ 1683ms | ✓ 1619ms | 否 | 否 | http |
| 86.104.72.220:1082 | 否 | ✓ 1207ms | ✓ 284ms | ✓ 1452ms | ✓ 1902ms | http |
| 160.238.65.8:3128 | ✓ 1029ms | ✓ 1821ms | ✓ 927ms | ✓ 1467ms | ✓ 1073ms | http |
| 160.238.65.3:3128 | ✓ 1029ms | ✓ 1844ms | ✓ 906ms | ✓ 1476ms | ✓ 1067ms | http |
| 160.238.65.5:3128 | ✓ 1030ms | ✓ 1829ms | ✓ 918ms | ✓ 1423ms | ✓ 1117ms | http |
| 103.157.200.126:3128 | ✓ 1218ms | 否 | ✓ 1343ms | ✓ 1921ms | ✓ 1399ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1463ms | ✓ 1383ms | ✓ 1312ms | http |
| 160.238.65.6:3128 | ✓ 1031ms | ✓ 1936ms | ✓ 987ms | 否 | ✓ 1292ms | http |
| 160.238.65.9:3128 | ✓ 1032ms | 否 | 否 | ✓ 1694ms | ✓ 1157ms | http |
| 103.35.190.69:1081 | ✓ 180ms | ✓ 1991ms | ✓ 427ms | 否 | 否 | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1462ms | ✓ 1130ms | ✓ 1713ms | http |
| 20.164.75.153:8080 | ✓ 1343ms | 否 | ✓ 1973ms | 否 | ✓ 1843ms | http |
| 120.92.212.16:8890 | ✓ 1676ms | ✓ 1619ms | 否 | ✓ 1707ms | ✓ 1899ms | http |
| 91.217.81.131:1080 | ✓ 753ms | 否 | ✓ 1374ms | 否 | ✓ 1942ms | http |
| 86.104.72.220:1081 | ✓ 1347ms | ✓ 1157ms | ✓ 144ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1009ms | ✓ 1677ms | ✓ 855ms | ✓ 1121ms | ✓ 898ms | http |
| 152.70.91.193:40000 | ✓ 1732ms | 否 | 否 | ✓ 1386ms | ✓ 1592ms | http |
| 103.35.190.69:1082 | ✓ 609ms | 否 | ✓ 1787ms | ✓ 986ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1603ms | ✓ 1596ms | ✓ 1122ms | http |
| 121.230.8.91:1080 | 否 | ✓ 1525ms | ✓ 1158ms | ✓ 1416ms | ✓ 1144ms | http |
| 43.133.44.89:8888 | ✓ 1725ms | ✓ 1243ms | 否 | 否 | ✓ 1765ms | http |
| 130.61.174.200:1080 | ✓ 1790ms | ✓ 1191ms | ✓ 750ms | 否 | ✓ 1444ms | http |
| 144.124.227.88:3128 | 否 | 否 | ✓ 1359ms | ✓ 1919ms | ✓ 1533ms | http |
| 220.197.44.36:3128 | ✓ 1544ms | ✓ 1812ms | ✓ 1663ms | 否 | 否 | http |
| 154.90.48.209:9090 | ✓ 931ms | 否 | ✓ 1190ms | ✓ 1567ms | ✓ 1111ms | http |
| 45.125.67.37:8443 | ✓ 1061ms | 否 | ✓ 1484ms | ✓ 1263ms | ✓ 1312ms | http |
| 152.42.177.32:8888 | ✓ 1096ms | 否 | ✓ 1212ms | ✓ 1325ms | ✓ 1343ms | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1398ms | ✓ 1560ms | ✓ 1368ms | http |
| 86.104.72.219:1082 | ✓ 1038ms | ✓ 1033ms | 否 | 否 | ✓ 1439ms | http |
| 92.119.127.208:6005 | ✓ 1108ms | ✓ 1573ms | ✓ 1863ms | ✓ 1994ms | 否 | http |
| 3.101.133.120:80 | ✓ 669ms | ✓ 1391ms | ✓ 1875ms | ✓ 1392ms | ✓ 1181ms | http |
| 94.131.118.129:1081 | ✓ 1050ms | ✓ 1305ms | ✓ 1095ms | 否 | ✓ 1213ms | http |
| 86.104.74.110:1081 | ✓ 1249ms | 否 | ✓ 465ms | 否 | ✓ 1271ms | http |
| 45.88.0.115:3128 | ✓ 740ms | ✓ 1709ms | ✓ 1317ms | 否 | ✓ 1443ms | http |
| 86.104.74.110:1082 | ✓ 1208ms | ✓ 1130ms | ✓ 1078ms | 否 | ✓ 1548ms | http |
| 187.102.219.32:999 | ✓ 1486ms | 否 | ✓ 1631ms | 否 | ✓ 1988ms | http |
| 45.88.0.117:3128 | ✓ 408ms | ✓ 1401ms | ✓ 747ms | ✓ 1242ms | ✓ 967ms | http |
| 45.88.0.111:3128 | ✓ 412ms | ✓ 1260ms | ✓ 880ms | ✓ 1242ms | ✓ 963ms | http |
| 213.220.62.63:3128 | ✓ 441ms | ✓ 1229ms | ✓ 882ms | ✓ 1242ms | ✓ 969ms | http |
| 45.88.0.116:3128 | ✓ 421ms | ✓ 1286ms | ✓ 850ms | ✓ 1244ms | ✓ 964ms | http |
| 45.88.0.98:3128 | ✓ 408ms | ✓ 1406ms | ✓ 743ms | ✓ 1265ms | ✓ 954ms | http |
| 45.88.0.113:3128 | ✓ 507ms | ✓ 1331ms | ✓ 707ms | ✓ 1254ms | ✓ 968ms | http |
| 38.180.62.47:10808 | ✓ 862ms | ✓ 1236ms | ✓ 918ms | ✓ 1985ms | 否 | http |
| 38.180.121.135:10808 | ✓ 585ms | 否 | ✓ 1563ms | ✓ 1872ms | ✓ 1549ms | http |
| 139.162.153.201:3128 | ✓ 773ms | 否 | ✓ 1664ms | 否 | ✓ 1890ms | http |
| 43.157.41.157:3128 | ✓ 772ms | ✓ 1857ms | ✓ 1615ms | 否 | ✓ 1417ms | http |
| 94.131.118.39:1081 | ✓ 886ms | ✓ 1492ms | ✓ 1180ms | ✓ 1456ms | ✓ 1480ms | http |
| 62.113.119.14:8080 | ✓ 1402ms | 否 | ✓ 565ms | ✓ 1548ms | ✓ 1075ms | http |
| 103.209.36.58:8080 | ✓ 1727ms | 否 | 否 | ✓ 1726ms | ✓ 1832ms | http |
| 94.131.118.39:1082 | ✓ 1140ms | 否 | ✓ 744ms | ✓ 1724ms | 否 | http |
| 149.202.47.125:3128 | ✓ 1121ms | ✓ 1770ms | ✓ 1465ms | 否 | 否 | http |
| 158.178.237.243:3128 | ✓ 1489ms | 否 | 否 | ✓ 1498ms | ✓ 1408ms | http |
| 210.223.44.230:3128 | ✓ 1720ms | ✓ 1510ms | ✓ 1902ms | ✓ 1174ms | ✓ 886ms | http |
| 148.153.56.51:80 | 否 | ✓ 979ms | ✓ 1084ms | ✓ 1070ms | ✓ 679ms | http |
| 77.110.107.80:1080 | ✓ 891ms | 否 | ✓ 1118ms | ✓ 1883ms | ✓ 1397ms | http |
| 77.110.107.80:8080 | ✓ 887ms | 否 | ✓ 1141ms | ✓ 1867ms | ✓ 1887ms | http |
| 116.171.106.111:3443 | 否 | 否 | ✓ 1619ms | ✓ 1534ms | ✓ 1657ms | http |
| 37.187.109.70:10111 | ✓ 511ms | 否 | ✓ 474ms | 否 | ✓ 1862ms | http |
| 61.52.131.172:8443 | ✓ 1000ms | ✓ 1301ms | ✓ 1039ms | ✓ 1429ms | ✓ 1082ms | http |
| 103.172.70.173:8080 | ✓ 1510ms | 否 | 否 | ✓ 1544ms | ✓ 1825ms | http |
| 121.230.8.136:1080 | ✓ 1058ms | ✓ 1467ms | ✓ 1245ms | ✓ 1660ms | ✓ 1205ms | http |
| 103.35.190.182:1081 | 否 | 否 | ✓ 1772ms | ✓ 1261ms | ✓ 763ms | http |
| 103.39.51.207:8080 | ✓ 1459ms | 否 | 否 | ✓ 1857ms | ✓ 1856ms | http |
| 168.222.254.136:8888 | ✓ 1191ms | ✓ 1763ms | 否 | 否 | ✓ 1660ms | http |

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
