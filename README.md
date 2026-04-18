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

最后更新：2026-04-18 15:36:43 UTC（2026-04-18 23:36:43 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.27.15.49:8561 | ✓ 623ms | ✓ 1646ms | ✓ 768ms | ✓ 1024ms | ✓ 935ms | http |
| 20.78.118.91:8561 | ✓ 619ms | 否 | ✓ 701ms | ✓ 1073ms | ✓ 778ms | http |
| 20.210.39.153:8561 | ✓ 622ms | ✓ 1901ms | ✓ 771ms | ✓ 1108ms | ✓ 776ms | http |
| 1.231.81.166:3128 | ✓ 940ms | ✓ 1516ms | ✓ 917ms | ✓ 1074ms | ✓ 814ms | http |
| 46.101.95.183:8888 | ✓ 972ms | 否 | ✓ 1331ms | ✓ 1541ms | ✓ 1303ms | http |
| 152.42.208.139:8118 | ✓ 838ms | 否 | ✓ 1390ms | ✓ 1194ms | ✓ 947ms | http |
| 194.104.9.38:3128 | ✓ 974ms | ✓ 1467ms | 否 | ✓ 1547ms | ✓ 1407ms | http |
| 81.30.156.115:8080 | ✓ 970ms | 否 | ✓ 1041ms | ✓ 1836ms | ✓ 1673ms | http |
| 113.160.132.26:8080 | ✓ 1600ms | ✓ 1540ms | ✓ 1103ms | ✓ 1459ms | ✓ 1549ms | http |
| 223.84.151.86:30005 | ✓ 1453ms | ✓ 1662ms | ✓ 1475ms | ✓ 1646ms | ✓ 1561ms | http |
| 14.247.76.52:8080 | ✓ 1604ms | 否 | ✓ 1040ms | ✓ 1819ms | ✓ 1183ms | http |
| 20.27.11.248:8561 | ✓ 685ms | ✓ 1623ms | 否 | 否 | ✓ 1868ms | http |
| 149.51.42.10:3128 | ✓ 562ms | ✓ 1940ms | 否 | ✓ 1369ms | 否 | http |
| 45.140.147.82:1081 | ✓ 412ms | ✓ 1218ms | ✓ 402ms | ✓ 1372ms | ✓ 1137ms | http |
| 185.138.116.150:8080 | ✓ 555ms | 否 | ✓ 1045ms | ✓ 1422ms | ✓ 1107ms | http |
| 133.18.123.225:26021 | ✓ 1017ms | 否 | ✓ 1157ms | ✓ 1270ms | ✓ 1252ms | http |
| 213.32.85.26:3128 | ✓ 532ms | 否 | ✓ 607ms | 否 | ✓ 1466ms | http |
| 116.58.161.203:26021 | ✓ 1097ms | 否 | ✓ 857ms | ✓ 1423ms | ✓ 1291ms | http |
| 188.246.224.49:7890 | ✓ 674ms | ✓ 1335ms | ✓ 1683ms | ✓ 1797ms | ✓ 1805ms | http |
| 177.93.132.244:3128 | ✓ 1081ms | 否 | ✓ 744ms | 否 | ✓ 1674ms | http |
| 94.131.118.129:1081 | ✓ 972ms | 否 | ✓ 1317ms | 否 | ✓ 1677ms | http |
| 45.67.202.178:1080 | ✓ 778ms | 否 | ✓ 1507ms | ✓ 1955ms | ✓ 1670ms | http |
| 162.19.253.202:8443 | ✓ 787ms | 否 | ✓ 1642ms | 否 | ✓ 1989ms | http |
| 84.47.150.126:1080 | ✓ 715ms | 否 | ✓ 1724ms | 否 | ✓ 1813ms | http |
| 84.47.150.125:1080 | ✓ 1320ms | 否 | 否 | ✓ 1963ms | ✓ 1643ms | http |
| 144.31.27.49:1080 | ✓ 729ms | ✓ 1653ms | 否 | 否 | ✓ 1524ms | http |
| 193.23.194.147:3128 | ✓ 735ms | ✓ 1943ms | ✓ 852ms | ✓ 1859ms | ✓ 1683ms | http |
| 20.78.26.206:8561 | ✓ 699ms | ✓ 1464ms | ✓ 633ms | ✓ 977ms | ✓ 744ms | http |
| 20.27.13.35:8561 | ✓ 701ms | ✓ 1091ms | ✓ 855ms | ✓ 996ms | ✓ 783ms | http |
| 20.27.14.220:8561 | ✓ 703ms | 否 | ✓ 619ms | ✓ 963ms | ✓ 772ms | http |
| 20.27.15.111:8561 | ✓ 701ms | 否 | ✓ 633ms | ✓ 957ms | ✓ 761ms | http |
| 20.210.76.104:8561 | ✓ 1018ms | ✓ 1281ms | ✓ 942ms | ✓ 1090ms | ✓ 949ms | http |
| 20.18.193.135:8561 | ✓ 1020ms | ✓ 1364ms | ✓ 870ms | ✓ 1071ms | ✓ 949ms | http |
| 20.210.76.178:8561 | ✓ 1024ms | ✓ 1863ms | ✓ 696ms | ✓ 1349ms | ✓ 904ms | http |
| 20.210.76.175:8561 | ✓ 1032ms | 否 | ✓ 629ms | ✓ 1267ms | ✓ 910ms | http |
| 91.99.15.45:2095 | ✓ 1092ms | ✓ 1433ms | ✓ 1362ms | ✓ 1970ms | ✓ 1806ms | http |
| 45.153.231.229:8080 | ✓ 1151ms | ✓ 1745ms | ✓ 1161ms | 否 | 否 | http |
| 149.51.42.10:8080 | ✓ 414ms | ✓ 1251ms | 否 | ✓ 1528ms | 否 | http |
| 218.108.131.186:17890 | ✓ 958ms | ✓ 1157ms | ✓ 1006ms | ✓ 1202ms | ✓ 1013ms | http |
| 89.111.174.221:8080 | ✓ 1849ms | 否 | ✓ 1698ms | 否 | ✓ 1743ms | http |
| 161.97.184.191:8080 | ✓ 1070ms | 否 | ✓ 921ms | 否 | ✓ 1794ms | http |
| 117.122.240.82:3338 | ✓ 1232ms | ✓ 1751ms | ✓ 1641ms | ✓ 1304ms | ✓ 1350ms | http |
| 120.92.212.16:7890 | ✓ 1071ms | ✓ 1625ms | 否 | ✓ 1845ms | ✓ 1145ms | http |
| 45.76.207.177:40000 | ✓ 790ms | 否 | ✓ 1211ms | ✓ 1161ms | ✓ 1231ms | http |
| 8.219.195.129:1080 | 否 | 否 | ✓ 992ms | ✓ 1175ms | ✓ 941ms | http |
| 35.225.22.61:80 | ✓ 745ms | ✓ 1287ms | 否 | 否 | ✓ 754ms | http |
| 62.113.119.14:8080 | ✓ 1463ms | ✓ 1763ms | ✓ 620ms | ✓ 1474ms | ✓ 1951ms | http |
| 117.236.124.166:3128 | ✓ 1126ms | 否 | ✓ 1134ms | 否 | ✓ 1226ms | http |
| 20.127.128.70:8080 | ✓ 1259ms | 否 | ✓ 1542ms | 否 | ✓ 1977ms | http |
| 45.140.147.82:1082 | ✓ 1717ms | 否 | ✓ 481ms | ✓ 1194ms | 否 | http |
| 178.63.155.151:9000 | ✓ 1059ms | 否 | ✓ 1368ms | 否 | ✓ 1738ms | http |
| 168.110.52.228:3128 | ✓ 762ms | 否 | ✓ 1789ms | ✓ 1135ms | ✓ 894ms | http |
| 106.10.55.212:1121 | ✓ 1275ms | ✓ 1331ms | ✓ 1417ms | ✓ 1538ms | ✓ 1293ms | http |
| 149.202.47.125:3128 | ✓ 1396ms | 否 | ✓ 1450ms | 否 | ✓ 1853ms | http |
| 116.236.189.92:29999 | 否 | ✓ 1275ms | ✓ 1076ms | 否 | ✓ 1068ms | http |
| 41.248.2.174:1221 | ✓ 1495ms | 否 | ✓ 1473ms | 否 | ✓ 1955ms | http |
| 43.132.188.134:443 | 否 | 否 | ✓ 805ms | ✓ 1021ms | ✓ 1416ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1143ms | ✓ 1988ms | ✓ 1778ms | ✓ 893ms | http |
| 160.238.65.6:3128 | ✓ 496ms | ✓ 1295ms | ✓ 1552ms | ✓ 1963ms | ✓ 1424ms | http |
| 160.238.65.2:3128 | ✓ 468ms | 否 | ✓ 849ms | ✓ 1961ms | ✓ 1431ms | http |
| 160.238.65.7:3128 | ✓ 471ms | ✓ 1326ms | ✓ 1512ms | ✓ 1969ms | ✓ 1430ms | http |
| 160.238.65.4:3128 | ✓ 496ms | 否 | ✓ 833ms | ✓ 1954ms | ✓ 1452ms | http |
| 160.238.65.5:3128 | ✓ 498ms | 否 | ✓ 826ms | ✓ 1985ms | ✓ 1418ms | http |
| 160.238.65.3:3128 | ✓ 470ms | ✓ 1339ms | ✓ 1505ms | ✓ 1963ms | ✓ 1425ms | http |
| 160.238.65.9:3128 | ✓ 470ms | 否 | ✓ 853ms | ✓ 1949ms | ✓ 1442ms | http |
| 162.240.154.26:3128 | ✓ 782ms | 否 | ✓ 1231ms | ✓ 1132ms | 否 | http |
| 138.124.99.216:8888 | ✓ 741ms | 否 | ✓ 1407ms | 否 | ✓ 1670ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 677ms | ✓ 1874ms | ✓ 1521ms | http |
| 101.32.243.189:80 | ✓ 1717ms | ✓ 1647ms | ✓ 1818ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 1102ms | ✓ 1588ms | ✓ 1264ms | 否 | 否 | http |
| 208.87.243.199:7878 | ✓ 1715ms | 否 | ✓ 1376ms | 否 | ✓ 1580ms | http |
| 147.161.166.35:10326 | ✓ 485ms | 否 | ✓ 1193ms | ✓ 1711ms | ✓ 1367ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1167ms | ✓ 1136ms | ✓ 1190ms | http |
| 103.113.70.189:1081 | ✓ 197ms | ✓ 1325ms | ✓ 480ms | ✓ 1053ms | ✓ 840ms | http |
| 103.113.70.189:1082 | ✓ 198ms | ✓ 1325ms | ✓ 482ms | ✓ 1061ms | ✓ 1133ms | http |
| 120.92.212.16:8890 | ✓ 1421ms | ✓ 1412ms | ✓ 1584ms | 否 | 否 | http |
| 152.32.132.190:7890 | 否 | ✓ 1965ms | 否 | ✓ 1254ms | ✓ 818ms | http |
| 195.86.215.2:3128 | ✓ 958ms | 否 | ✓ 988ms | ✓ 1240ms | ✓ 977ms | http |
| 194.87.130.16:1080 | ✓ 1485ms | 否 | ✓ 1408ms | 否 | ✓ 1453ms | http |
| 111.79.111.126:3128 | ✓ 1211ms | ✓ 1644ms | ✓ 1976ms | 否 | ✓ 1260ms | http |
| 45.140.147.155:1081 | ✓ 1957ms | ✓ 1255ms | ✓ 845ms | ✓ 1583ms | ✓ 1328ms | http |
| 104.247.51.76:3128 | 否 | 否 | ✓ 991ms | ✓ 1080ms | ✓ 836ms | http |
| 157.0.142.246:10061 | ✓ 1143ms | ✓ 1429ms | ✓ 1162ms | ✓ 1453ms | ✓ 1180ms | http |
| 147.45.166.46:3128 | 否 | 否 | ✓ 1937ms | ✓ 1608ms | ✓ 1624ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 658ms | ✓ 963ms | ✓ 1204ms | http |
| 20.120.225.109:3128 | ✓ 714ms | ✓ 1314ms | ✓ 1869ms | 否 | 否 | http |
| 130.61.30.221:8080 | ✓ 923ms | ✓ 1959ms | ✓ 690ms | ✓ 1915ms | ✓ 1949ms | http |
| 103.85.113.66:9999 | ✓ 687ms | ✓ 1461ms | ✓ 1091ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 972ms | ✓ 1264ms | ✓ 1051ms | ✓ 1293ms | ✓ 1045ms | http |
| 103.135.102.161:8080 | ✓ 1765ms | 否 | ✓ 890ms | ✓ 924ms | ✓ 1760ms | http |
| 212.58.132.5:8888 | ✓ 1113ms | 否 | ✓ 1650ms | ✓ 1800ms | ✓ 1169ms | http |
| 103.133.254.4:3128 | ✓ 1783ms | 否 | ✓ 1634ms | ✓ 1708ms | ✓ 1360ms | http |
| 160.250.5.22:1 | ✓ 1562ms | 否 | ✓ 1862ms | ✓ 1744ms | ✓ 1304ms | http |

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
