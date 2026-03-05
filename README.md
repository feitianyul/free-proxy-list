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

最后更新：2026-03-05 08:49:05 UTC（2026-03-05 16:49:05 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 61.72.110.94:3128 | ✓ 1217ms | ✓ 1293ms | ✓ 1174ms | ✓ 966ms | ✓ 814ms | http |
| 205.209.118.30:3138 | ✓ 848ms | 否 | ✓ 1278ms | ✓ 1391ms | ✓ 1028ms | http |
| 125.128.12.14:3128 | ✓ 1221ms | ✓ 1883ms | ✓ 1170ms | ✓ 931ms | ✓ 804ms | http |
| 125.128.12.144:3128 | ✓ 1229ms | ✓ 1358ms | ✓ 1533ms | ✓ 994ms | ✓ 960ms | http |
| 61.72.221.234:3128 | ✓ 1222ms | ✓ 1576ms | 否 | 否 | ✓ 1785ms | http |
| 185.191.236.162:3128 | ✓ 1654ms | 否 | ✓ 1777ms | 否 | ✓ 1763ms | http |
| 61.72.221.194:3128 | ✓ 1218ms | 否 | ✓ 1967ms | 否 | ✓ 1168ms | http |
| 61.72.221.94:3128 | ✓ 1084ms | ✓ 1093ms | ✓ 1654ms | ✓ 1317ms | 否 | http |
| 61.72.110.54:3128 | ✓ 607ms | ✓ 1041ms | 否 | ✓ 995ms | ✓ 944ms | http |
| 121.128.121.54:3128 | ✓ 529ms | ✓ 1811ms | 否 | 否 | ✓ 1728ms | http |
| 14.56.107.244:3128 | ✓ 536ms | 否 | ✓ 562ms | ✓ 1938ms | 否 | http |
| 217.76.245.80:999 | ✓ 997ms | 否 | ✓ 1297ms | ✓ 1658ms | ✓ 1405ms | http |
| 116.80.82.232:3172 | ✓ 1867ms | 否 | ✓ 1448ms | ✓ 1794ms | ✓ 1868ms | http |
| 116.80.82.220:3172 | 否 | 否 | ✓ 1549ms | ✓ 1802ms | ✓ 1646ms | http |
| 116.80.82.228:3172 | ✓ 1868ms | 否 | ✓ 1452ms | 否 | ✓ 1796ms | http |
| 116.80.82.224:3172 | ✓ 1868ms | 否 | ✓ 1671ms | 否 | ✓ 1625ms | http |
| 103.84.95.54:7890 | ✓ 598ms | 否 | ✓ 869ms | ✓ 1288ms | ✓ 612ms | http |
| 120.92.212.16:8890 | ✓ 926ms | ✓ 1954ms | ✓ 1215ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 937ms | ✓ 1175ms | ✓ 1154ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1571ms | ✓ 1485ms | ✓ 845ms | 否 | ✓ 681ms | http |
| 168.235.110.63:3128 | ✓ 326ms | ✓ 1304ms | ✓ 1296ms | ✓ 1405ms | ✓ 1016ms | http |
| 101.43.255.96:80 | ✓ 907ms | ✓ 1625ms | 否 | ✓ 1190ms | ✓ 1967ms | http |
| 39.104.201.40:7890 | ✓ 1925ms | ✓ 1235ms | ✓ 1835ms | 否 | 否 | http |
| 116.80.82.222:3172 | ✓ 1764ms | 否 | ✓ 1539ms | ✓ 1832ms | ✓ 1594ms | http |
| 116.80.82.227:3172 | ✓ 1764ms | 否 | ✓ 1539ms | ✓ 1794ms | ✓ 1634ms | http |
| 116.80.82.226:3172 | ✓ 1764ms | 否 | ✓ 1539ms | ✓ 1799ms | ✓ 1639ms | http |
| 116.80.82.230:3172 | ✓ 1764ms | 否 | ✓ 1683ms | ✓ 1781ms | 否 | http |
| 116.80.82.223:3172 | ✓ 1764ms | 否 | ✓ 1539ms | ✓ 1816ms | ✓ 1609ms | http |
| 120.92.212.16:7890 | ✓ 1176ms | ✓ 1194ms | ✓ 1169ms | ✓ 1207ms | ✓ 1150ms | http |
| 95.85.252.153:21064 | ✓ 780ms | 否 | ✓ 1048ms | 否 | ✓ 1641ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1391ms | ✓ 1764ms | ✓ 1257ms | ✓ 1442ms | http |
| 62.113.119.14:8080 | ✓ 771ms | 否 | ✓ 1429ms | ✓ 1658ms | ✓ 1246ms | http |
| 91.193.240.157:9877 | ✓ 891ms | 否 | ✓ 1034ms | 否 | ✓ 1589ms | http |
| 121.204.158.249:3128 | ✓ 1084ms | ✓ 1271ms | ✓ 1236ms | ✓ 1343ms | ✓ 1383ms | http |
| 222.228.171.92:8080 | ✓ 1593ms | 否 | ✓ 806ms | 否 | ✓ 1404ms | http |
| 94.176.3.43:7443 | ✓ 1777ms | 否 | ✓ 1268ms | ✓ 1638ms | 否 | http |
| 116.80.82.231:3172 | ✓ 1565ms | 否 | ✓ 1514ms | ✓ 1798ms | 否 | http |
| 116.80.82.221:3172 | ✓ 1870ms | 否 | ✓ 1463ms | ✓ 1821ms | ✓ 1761ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1876ms | 否 | ✓ 1296ms | ✓ 981ms | http |
| 121.237.181.137:8888 | ✓ 926ms | ✓ 1031ms | 否 | ✓ 1954ms | ✓ 929ms | http |
| 160.250.5.22:1 | ✓ 1914ms | 否 | ✓ 1398ms | ✓ 1374ms | ✓ 1072ms | http |
| 81.70.169.194:80 | ✓ 1267ms | ✓ 1771ms | ✓ 993ms | ✓ 1193ms | ✓ 1221ms | http |
| 111.79.111.126:3128 | ✓ 1849ms | ✓ 1571ms | ✓ 1601ms | 否 | ✓ 1280ms | http |
| 154.236.177.120:1981 | 否 | 否 | ✓ 1449ms | ✓ 1873ms | ✓ 1711ms | http |
| 45.140.147.82:1081 | ✓ 604ms | ✓ 1857ms | ✓ 1459ms | ✓ 1582ms | ✓ 1272ms | http |
| 58.69.222.240:8082 | ✓ 1892ms | 否 | ✓ 1670ms | ✓ 1379ms | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1271ms | 否 | ✓ 1030ms | ✓ 993ms | http |
| 188.132.141.249:443 | ✓ 1678ms | 否 | ✓ 1294ms | 否 | ✓ 1755ms | http |
| 45.140.147.82:1082 | ✓ 643ms | 否 | ✓ 1039ms | 否 | ✓ 1637ms | http |
| 45.136.198.40:3128 | ✓ 948ms | 否 | ✓ 1360ms | 否 | ✓ 1730ms | http |
| 58.220.95.11:11023 | ✓ 1692ms | ✓ 1136ms | ✓ 1163ms | ✓ 1217ms | ✓ 860ms | http |
| 183.237.195.130:3128 | ✓ 896ms | ✓ 1211ms | ✓ 1077ms | ✓ 1153ms | ✓ 1925ms | http |
| 121.230.9.161:1080 | ✓ 1091ms | 否 | 否 | ✓ 1342ms | ✓ 1170ms | http |
| 185.243.218.43:49153 | ✓ 1808ms | 否 | ✓ 1752ms | ✓ 1966ms | ✓ 1999ms | http |
| 101.32.244.83:8080 | ✓ 1655ms | ✓ 1657ms | ✓ 899ms | ✓ 1108ms | ✓ 1169ms | http |
| 121.43.196.213:8222 | ✓ 976ms | ✓ 997ms | ✓ 907ms | ✓ 1072ms | ✓ 829ms | http |
| 121.43.196.210:8222 | ✓ 968ms | ✓ 1016ms | ✓ 895ms | ✓ 1073ms | ✓ 828ms | http |
| 114.55.226.123:10086 | ✓ 1082ms | ✓ 1651ms | ✓ 1017ms | ✓ 1198ms | ✓ 958ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 897ms | ✓ 1092ms | ✓ 849ms | http |
| 165.227.5.10:8888 | ✓ 1229ms | 否 | ✓ 229ms | ✓ 785ms | 否 | http |
| 193.228.139.78:8888 | ✓ 1546ms | 否 | ✓ 1705ms | 否 | ✓ 1921ms | http |
| 5.75.196.26:40000 | ✓ 1862ms | 否 | ✓ 1670ms | 否 | ✓ 1943ms | http |
| 121.230.8.136:1080 | ✓ 1688ms | ✓ 1280ms | ✓ 1934ms | ✓ 1666ms | ✓ 1182ms | http |
| 90.84.188.97:8000 | ✓ 1394ms | ✓ 1994ms | 否 | 否 | ✓ 1963ms | http |
| 160.238.65.8:3128 | ✓ 1359ms | 否 | ✓ 1428ms | 否 | ✓ 1840ms | http |
| 160.238.65.4:3128 | ✓ 887ms | ✓ 1689ms | ✓ 813ms | 否 | ✓ 1870ms | http |
| 160.238.65.2:3128 | ✓ 1328ms | 否 | ✓ 1430ms | 否 | ✓ 1865ms | http |
| 160.238.65.6:3128 | ✓ 1363ms | 否 | ✓ 1430ms | 否 | ✓ 1850ms | http |
| 160.238.65.3:3128 | ✓ 1161ms | ✓ 1860ms | ✓ 1565ms | 否 | ✓ 1872ms | http |
| 160.238.65.7:3128 | ✓ 1339ms | 否 | ✓ 1423ms | 否 | ✓ 1847ms | http |
| 160.238.65.5:3128 | ✓ 1353ms | 否 | ✓ 1434ms | 否 | ✓ 1859ms | http |
| 160.238.65.9:3128 | ✓ 617ms | 否 | ✓ 641ms | ✓ 1507ms | ✓ 1153ms | http |
| 103.215.36.88:17105 | ✓ 1012ms | 否 | ✓ 1095ms | 否 | ✓ 1248ms | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1914ms | ✓ 1729ms | ✓ 1535ms | http |
| 182.253.160.232:1452 | 否 | 否 | ✓ 1879ms | ✓ 1324ms | ✓ 1473ms | http |
| 47.74.226.8:5001 | ✓ 1193ms | ✓ 1473ms | ✓ 1639ms | 否 | 否 | http |
| 116.80.82.225:3172 | 否 | 否 | ✓ 1886ms | ✓ 1791ms | ✓ 1622ms | http |
| 202.129.206.239:3128 | ✓ 1639ms | 否 | 否 | ✓ 1426ms | ✓ 1296ms | http |
| 64.181.240.152:3128 | ✓ 1520ms | ✓ 1886ms | ✓ 513ms | ✓ 876ms | ✓ 587ms | http |
| 103.39.51.190:8080 | ✓ 1988ms | 否 | 否 | ✓ 1902ms | ✓ 1475ms | http |
| 121.230.9.252:1080 | ✓ 1566ms | ✓ 1669ms | ✓ 1904ms | ✓ 1377ms | ✓ 1152ms | http |
| 103.102.12.67:8080 | ✓ 1946ms | 否 | 否 | ✓ 1468ms | ✓ 1332ms | http |
| 8.137.112.117:3128 | ✓ 947ms | 否 | ✓ 1007ms | ✓ 1220ms | 否 | http |

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
