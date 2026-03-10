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

最后更新：2026-03-10 23:23:31 UTC（2026-03-11 07:23:31 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 905ms | ✓ 1828ms | ✓ 343ms | ✓ 980ms | ✓ 766ms | http |
| 45.136.130.175:8443 | ✓ 1120ms | ✓ 1101ms | ✓ 873ms | ✓ 968ms | ✓ 765ms | http |
| 45.136.130.191:8443 | ✓ 906ms | ✓ 1544ms | ✓ 1254ms | ✓ 963ms | ✓ 750ms | http |
| 1.231.81.166:3128 | ✓ 1525ms | ✓ 1147ms | ✓ 1243ms | ✓ 1233ms | ✓ 1066ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1686ms | ✓ 1405ms | ✓ 1634ms | ✓ 1608ms | http |
| 185.115.74.185:8080 | ✓ 1228ms | ✓ 1912ms | ✓ 1805ms | 否 | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1366ms | ✓ 1091ms | ✓ 1420ms | ✓ 1033ms | http |
| 190.9.109.198:999 | ✓ 749ms | ✓ 1390ms | ✓ 1095ms | ✓ 1452ms | ✓ 1045ms | http |
| 217.77.102.18:3128 | ✓ 1007ms | 否 | ✓ 1216ms | ✓ 1755ms | ✓ 1385ms | http |
| 103.137.193.55:8888 | ✓ 868ms | 否 | ✓ 1100ms | ✓ 1992ms | ✓ 1459ms | http |
| 35.225.22.61:80 | ✓ 332ms | ✓ 1231ms | 否 | 否 | ✓ 934ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1549ms | 否 | ✓ 1615ms | ✓ 1271ms | http |
| 46.183.25.8:443 | ✓ 1616ms | 否 | ✓ 695ms | ✓ 1294ms | ✓ 984ms | http |
| 202.155.12.161:443 | ✓ 1887ms | 否 | ✓ 1395ms | ✓ 1357ms | ✓ 1166ms | http |
| 158.69.185.37:3129 | ✓ 938ms | 否 | ✓ 1100ms | ✓ 1207ms | ✓ 844ms | http |
| 67.169.98.211:443 | ✓ 935ms | 否 | ✓ 649ms | ✓ 1769ms | 否 | http |
| 45.136.131.47:8443 | 否 | ✓ 907ms | ✓ 958ms | ✓ 1073ms | ✓ 779ms | http |
| 39.104.201.40:7890 | ✓ 1087ms | ✓ 1424ms | ✓ 1086ms | ✓ 1400ms | ✓ 1138ms | http |
| 101.43.255.96:80 | ✓ 1126ms | ✓ 1474ms | ✓ 1209ms | ✓ 1516ms | ✓ 1143ms | http |
| 81.70.169.194:80 | ✓ 1207ms | ✓ 1491ms | ✓ 1119ms | ✓ 1369ms | ✓ 1247ms | http |
| 45.136.130.188:8443 | ✓ 921ms | ✓ 905ms | ✓ 1543ms | ✓ 990ms | ✓ 737ms | http |
| 45.136.130.223:8443 | ✓ 918ms | ✓ 945ms | ✓ 1502ms | ✓ 995ms | ✓ 753ms | http |
| 152.70.98.46:8888 | ✓ 1139ms | ✓ 1220ms | ✓ 689ms | ✓ 986ms | ✓ 789ms | http |
| 45.136.130.239:8443 | 否 | ✓ 947ms | 否 | ✓ 981ms | ✓ 743ms | http |
| 152.42.213.210:8080 | ✓ 969ms | 否 | ✓ 1719ms | ✓ 1303ms | ✓ 1023ms | http |
| 194.213.18.200:443 | ✓ 1637ms | 否 | ✓ 1839ms | 否 | ✓ 1974ms | http |
| 120.92.212.16:7890 | ✓ 1132ms | ✓ 1415ms | ✓ 1423ms | ✓ 1716ms | ✓ 1420ms | http |
| 120.92.212.16:8890 | ✓ 1375ms | ✓ 1428ms | ✓ 1133ms | ✓ 1406ms | ✓ 1113ms | http |
| 138.124.53.25:7443 | ✓ 367ms | 否 | ✓ 1473ms | 否 | ✓ 1263ms | http |
| 86.53.183.16:1080 | ✓ 1194ms | ✓ 1595ms | ✓ 1246ms | 否 | 否 | http |
| 91.107.141.42:8081 | ✓ 1754ms | 否 | ✓ 1298ms | 否 | ✓ 1501ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1948ms | ✓ 1826ms | ✓ 1253ms | http |
| 168.235.110.63:3128 | ✓ 348ms | 否 | ✓ 437ms | ✓ 1073ms | ✓ 810ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1964ms | ✓ 1738ms | ✓ 1333ms | http |
| 62.113.119.14:8080 | ✓ 502ms | ✓ 1394ms | ✓ 1205ms | ✓ 1409ms | ✓ 1101ms | http |
| 72.56.104.188:1080 | ✓ 657ms | ✓ 1933ms | ✓ 1457ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 965ms | ✓ 1939ms | ✓ 375ms | ✓ 1308ms | ✓ 844ms | http |
| 94.176.3.43:7443 | ✓ 1195ms | 否 | ✓ 1108ms | 否 | ✓ 1419ms | http |
| 113.177.131.2:3128 | ✓ 1753ms | 否 | ✓ 1356ms | ✓ 1744ms | ✓ 1429ms | http |
| 222.228.171.92:8080 | ✓ 1630ms | 否 | ✓ 977ms | 否 | ✓ 1132ms | http |
| 121.230.8.26:1080 | ✓ 1227ms | ✓ 1806ms | ✓ 1225ms | ✓ 1891ms | ✓ 1205ms | http |
| 121.230.8.22:1080 | ✓ 1355ms | ✓ 1479ms | ✓ 1238ms | 否 | 否 | http |
| 47.77.193.180:1080 | ✓ 871ms | ✓ 977ms | ✓ 477ms | ✓ 943ms | ✓ 701ms | http |
| 34.96.238.40:8080 | ✓ 1353ms | 否 | ✓ 1317ms | ✓ 1209ms | 否 | http |
| 120.238.159.234:22222 | ✓ 1043ms | ✓ 1439ms | ✓ 1173ms | ✓ 1343ms | ✓ 1032ms | http |
| 45.77.246.231:80 | 否 | 否 | ✓ 1344ms | ✓ 1305ms | ✓ 1028ms | http |
| 205.209.118.30:3138 | ✓ 630ms | ✓ 890ms | ✓ 626ms | ✓ 1637ms | ✓ 838ms | http |
| 45.129.141.143:3128 | ✓ 1032ms | ✓ 1984ms | ✓ 1477ms | ✓ 1981ms | ✓ 1636ms | http |
| 95.3.9.78:3128 | ✓ 1350ms | 否 | 否 | ✓ 1594ms | ✓ 1204ms | http |
| 185.191.236.162:3128 | ✓ 1348ms | ✓ 1904ms | ✓ 1655ms | 否 | ✓ 1813ms | http |
| 45.186.6.104:3128 | ✓ 1151ms | ✓ 1793ms | ✓ 1626ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1991ms | 否 | 否 | ✓ 1682ms | ✓ 1981ms | http |
| 103.35.188.243:3128 | ✓ 120ms | ✓ 1122ms | 否 | ✓ 1076ms | ✓ 811ms | http |
| 45.136.198.40:3128 | ✓ 1081ms | 否 | ✓ 1599ms | 否 | ✓ 1495ms | http |
| 200.174.198.32:8888 | ✓ 1047ms | 否 | ✓ 1596ms | 否 | ✓ 1923ms | http |
| 121.230.9.160:1080 | 否 | ✓ 1748ms | ✓ 1334ms | ✓ 1590ms | ✓ 1292ms | http |
| 162.240.154.26:3128 | ✓ 1061ms | ✓ 1221ms | ✓ 1058ms | ✓ 1249ms | ✓ 1016ms | http |
| 138.124.90.140:1080 | ✓ 1174ms | ✓ 1443ms | ✓ 1539ms | ✓ 1542ms | 否 | http |
| 95.3.9.78:8080 | ✓ 904ms | ✓ 1707ms | ✓ 650ms | ✓ 1600ms | ✓ 1225ms | http |
| 45.140.147.82:1081 | ✓ 385ms | ✓ 1271ms | ✓ 939ms | ✓ 1688ms | ✓ 1264ms | http |
| 211.171.114.154:3128 | ✓ 1962ms | 否 | 否 | ✓ 1932ms | ✓ 1990ms | http |
| 121.230.8.55:1080 | ✓ 1282ms | ✓ 1593ms | ✓ 1354ms | 否 | ✓ 1196ms | http |
| 61.52.131.172:8443 | ✓ 1006ms | ✓ 1328ms | ✓ 1088ms | ✓ 1379ms | ✓ 1049ms | http |
| 166.249.54.61:7234 | ✓ 907ms | ✓ 1863ms | 否 | ✓ 1450ms | ✓ 1173ms | http |
| 202.129.206.239:3128 | ✓ 1628ms | 否 | ✓ 1721ms | 否 | ✓ 1490ms | http |
| 190.212.131.238:3128 | ✓ 1678ms | 否 | ✓ 1163ms | 否 | ✓ 1782ms | http |
| 5.9.55.221:5000 | ✓ 1074ms | ✓ 1685ms | 否 | 否 | ✓ 1766ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1828ms | ✓ 1834ms | ✓ 1292ms | ✓ 1059ms | http |
| 34.101.184.164:3128 | ✓ 1888ms | 否 | ✓ 1640ms | ✓ 1529ms | ✓ 1429ms | http |

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
