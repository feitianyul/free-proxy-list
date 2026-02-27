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

最后更新：2026-02-27 11:40:03 UTC（2026-02-27 19:40:03 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 128ms | 否 | ✓ 1117ms | ✓ 1309ms | ✓ 806ms | http |
| 103.113.70.189:1081 | 否 | ✓ 933ms | 否 | ✓ 1135ms | ✓ 808ms | http |
| 120.92.212.16:8890 | ✓ 1186ms | ✓ 1490ms | 否 | 否 | ✓ 1161ms | http |
| 147.45.216.148:1080 | ✓ 522ms | 否 | ✓ 1225ms | ✓ 1751ms | ✓ 1550ms | http |
| 120.240.35.173:22222 | ✓ 1108ms | 否 | ✓ 1186ms | 否 | ✓ 1103ms | http |
| 183.249.5.109:22222 | ✓ 917ms | ✓ 1418ms | ✓ 1107ms | 否 | 否 | http |
| 104.37.184.214:1080 | ✓ 1518ms | ✓ 1936ms | 否 | 否 | ✓ 1997ms | http |
| 195.123.209.48:3128 | ✓ 947ms | ✓ 1751ms | ✓ 1242ms | ✓ 1909ms | ✓ 1499ms | http |
| 223.113.134.102:22222 | ✓ 852ms | ✓ 1333ms | 否 | ✓ 1172ms | ✓ 892ms | http |
| 113.59.32.162:22222 | 否 | ✓ 1552ms | 否 | ✓ 1412ms | ✓ 1233ms | http |
| 85.208.108.43:10808 | ✓ 1931ms | 否 | ✓ 1004ms | ✓ 1347ms | ✓ 922ms | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 1158ms | ✓ 1278ms | ✓ 919ms | http |
| 85.208.108.43:2094 | ✓ 452ms | 否 | ✓ 631ms | ✓ 1153ms | ✓ 890ms | http |
| 185.246.90.163:10808 | ✓ 655ms | 否 | ✓ 1360ms | ✓ 1775ms | ✓ 1714ms | http |
| 104.238.30.91:63900 | ✓ 1566ms | 否 | ✓ 1967ms | 否 | ✓ 1938ms | http |
| 45.140.147.82:1082 | ✓ 831ms | 否 | ✓ 1014ms | 否 | ✓ 1144ms | http |
| 104.238.30.63:63744 | ✓ 1598ms | 否 | ✓ 1683ms | 否 | ✓ 1931ms | http |
| 72.56.59.62:63133 | ✓ 1431ms | 否 | ✓ 1523ms | 否 | ✓ 1803ms | http |
| 168.235.110.63:3128 | ✓ 371ms | ✓ 995ms | ✓ 1104ms | ✓ 1161ms | ✓ 744ms | http |
| 223.113.134.105:22222 | ✓ 957ms | ✓ 1087ms | ✓ 1258ms | 否 | ✓ 1067ms | http |
| 223.113.134.141:22222 | ✓ 906ms | ✓ 1343ms | ✓ 929ms | ✓ 1179ms | ✓ 854ms | http |
| 81.70.169.194:80 | 否 | ✓ 1518ms | ✓ 1359ms | 否 | ✓ 1225ms | http |
| 101.43.255.96:80 | 否 | ✓ 1926ms | ✓ 1167ms | ✓ 1550ms | ✓ 1237ms | http |
| 104.238.30.58:63744 | ✓ 1595ms | 否 | ✓ 1903ms | 否 | ✓ 1939ms | http |
| 120.232.242.119:22222 | ✓ 1228ms | ✓ 1445ms | 否 | ✓ 1345ms | 否 | http |
| 72.56.59.56:63127 | ✓ 1438ms | 否 | ✓ 1523ms | 否 | ✓ 1738ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1709ms | 否 | ✓ 1284ms | ✓ 1045ms | http |
| 104.238.30.86:63900 | ✓ 1613ms | 否 | ✓ 1743ms | 否 | ✓ 1935ms | http |
| 35.234.17.221:8080 | ✓ 1070ms | 否 | 否 | ✓ 1512ms | ✓ 1091ms | http |
| 52.188.28.218:3128 | ✓ 81ms | ✓ 1026ms | ✓ 46ms | ✓ 905ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1783ms | 否 | ✓ 1826ms | ✓ 1814ms | 否 | http |
| 35.225.22.61:80 | ✓ 974ms | ✓ 1949ms | ✓ 1426ms | ✓ 1125ms | ✓ 869ms | http |
| 94.177.131.12:3128 | 否 | 否 | ✓ 1008ms | ✓ 1379ms | ✓ 1080ms | http |
| 45.88.0.117:3128 | ✓ 1630ms | 否 | 否 | ✓ 1219ms | ✓ 1256ms | http |
| 185.243.218.43:49153 | ✓ 516ms | 否 | ✓ 1357ms | ✓ 1765ms | ✓ 1498ms | http |
| 81.177.48.54:2080 | ✓ 692ms | 否 | ✓ 1394ms | 否 | ✓ 1294ms | http |
| 72.56.59.17:61931 | ✓ 1427ms | 否 | ✓ 1523ms | 否 | ✓ 1743ms | http |
| 72.56.59.23:61937 | ✓ 1461ms | 否 | ✓ 1522ms | 否 | ✓ 1768ms | http |
| 104.238.30.45:59741 | ✓ 1593ms | 否 | ✓ 1679ms | 否 | ✓ 1935ms | http |
| 104.238.30.40:59741 | ✓ 1587ms | 否 | ✓ 1679ms | 否 | ✓ 1935ms | http |
| 34.142.0.1:10808 | ✓ 518ms | 否 | ✓ 1168ms | ✓ 1899ms | ✓ 1614ms | http |
| 104.238.30.50:59741 | ✓ 1589ms | 否 | ✓ 1643ms | 否 | ✓ 1903ms | http |
| 37.27.100.108:443 | 否 | ✓ 1621ms | ✓ 848ms | 否 | ✓ 1940ms | http |
| 104.238.30.37:59741 | ✓ 1607ms | 否 | ✓ 1739ms | 否 | ✓ 1903ms | http |
| 115.231.181.40:8128 | ✓ 1143ms | ✓ 1450ms | ✓ 1075ms | 否 | 否 | http |
| 104.238.30.39:59741 | ✓ 1589ms | 否 | ✓ 1649ms | 否 | ✓ 1969ms | http |
| 138.124.53.25:7443 | ✓ 977ms | 否 | 否 | ✓ 1968ms | ✓ 1777ms | http |
| 104.238.30.38:59741 | ✓ 1591ms | 否 | ✓ 1679ms | 否 | ✓ 1940ms | http |
| 113.59.32.160:22222 | ✓ 1139ms | ✓ 1442ms | ✓ 1043ms | ✓ 1451ms | ✓ 1616ms | http |
| 223.113.134.103:22222 | ✓ 872ms | ✓ 1131ms | ✓ 852ms | ✓ 1182ms | ✓ 889ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1781ms | 否 | ✓ 1560ms | ✓ 1268ms | http |
| 34.101.184.164:3128 | ✓ 1261ms | 否 | ✓ 1742ms | ✓ 1732ms | 否 | http |
| 72.56.50.17:59787 | ✓ 1450ms | 否 | ✓ 1455ms | 否 | ✓ 1747ms | http |
| 223.113.134.57:22222 | ✓ 953ms | ✓ 1418ms | ✓ 944ms | ✓ 1122ms | ✓ 961ms | http |
| 223.113.134.104:22222 | 否 | ✓ 1080ms | ✓ 912ms | ✓ 1133ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1207ms | ✓ 1753ms | 否 | 否 | ✓ 1177ms | http |
| 120.238.159.229:22222 | 否 | ✓ 1507ms | 否 | ✓ 1429ms | ✓ 1073ms | http |
| 103.117.100.127:13082 | ✓ 827ms | 否 | ✓ 816ms | ✓ 1060ms | ✓ 824ms | http |
| 120.198.141.75:22222 | ✓ 1077ms | 否 | ✓ 1144ms | ✓ 1383ms | 否 | http |
| 113.59.32.148:22222 | ✓ 1279ms | ✓ 1596ms | ✓ 1768ms | 否 | 否 | http |
| 120.240.35.176:22222 | ✓ 1121ms | ✓ 1495ms | 否 | ✓ 1382ms | ✓ 1090ms | http |
| 113.59.32.141:22222 | 否 | ✓ 1555ms | ✓ 1071ms | ✓ 1471ms | ✓ 1081ms | http |
| 120.132.97.88:7897 | ✓ 1395ms | ✓ 1460ms | ✓ 1845ms | ✓ 1521ms | ✓ 1040ms | http |
| 103.215.36.88:15247 | ✓ 1269ms | ✓ 1791ms | 否 | ✓ 1911ms | ✓ 1269ms | http |
| 147.45.251.242:8888 | ✓ 1229ms | ✓ 1993ms | ✓ 1437ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1183ms | 否 | ✓ 1051ms | ✓ 1534ms | ✓ 1144ms | http |
| 45.125.67.37:8443 | ✓ 1265ms | 否 | 否 | ✓ 1552ms | ✓ 1102ms | http |
| 113.59.32.161:22222 | ✓ 1170ms | ✓ 1470ms | ✓ 1196ms | ✓ 1401ms | ✓ 1139ms | http |
| 59.46.216.131:30001 | ✓ 1203ms | 否 | ✓ 1295ms | 否 | ✓ 1250ms | http |
| 120.46.152.136:3128 | ✓ 1181ms | ✓ 1229ms | 否 | ✓ 1803ms | ✓ 1866ms | http |
| 54.88.116.133:80 | ✓ 175ms | ✓ 1435ms | ✓ 736ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1997ms | ✓ 1813ms | ✓ 1410ms | 否 | ✓ 1473ms | http |
| 103.84.95.54:7890 | ✓ 1057ms | 否 | 否 | ✓ 1045ms | ✓ 827ms | http |
| 14.56.107.244:3128 | ✓ 1985ms | ✓ 1395ms | ✓ 1009ms | 否 | ✓ 1422ms | http |
| 54.161.78.88:80 | ✓ 294ms | ✓ 868ms | ✓ 1108ms | ✓ 1423ms | ✓ 1133ms | http |
| 50.19.255.0:80 | ✓ 296ms | 否 | ✓ 52ms | ✓ 1374ms | ✓ 1111ms | http |
| 3.225.78.45:80 | ✓ 329ms | 否 | ✓ 1104ms | 否 | ✓ 1157ms | http |
| 45.88.0.98:3128 | ✓ 900ms | 否 | ✓ 704ms | 否 | ✓ 1025ms | http |
| 47.106.73.57:8118 | ✓ 1765ms | 否 | ✓ 1780ms | 否 | ✓ 1416ms | http |
| 44.219.16.212:80 | ✓ 789ms | ✓ 1730ms | ✓ 766ms | ✓ 1345ms | ✓ 1543ms | http |
| 223.113.134.92:22222 | ✓ 868ms | ✓ 1494ms | 否 | ✓ 1224ms | ✓ 879ms | http |
| 117.159.239.54:22222 | ✓ 1028ms | ✓ 1358ms | ✓ 1114ms | ✓ 1361ms | 否 | http |
| 100.50.110.44:80 | ✓ 788ms | ✓ 1420ms | ✓ 891ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1572ms | 否 | ✓ 1142ms | ✓ 1451ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1572ms | 否 | 否 | ✓ 1775ms | ✓ 1866ms | http |
| 3.214.214.245:80 | ✓ 923ms | ✓ 1220ms | ✓ 748ms | ✓ 1141ms | 否 | http |
| 144.124.227.88:3128 | ✓ 862ms | 否 | ✓ 999ms | ✓ 1861ms | 否 | http |
| 120.198.141.79:22222 | 否 | 否 | ✓ 1382ms | ✓ 1410ms | ✓ 1105ms | http |
| 36.147.78.166:443 | 否 | ✓ 1910ms | ✓ 1967ms | 否 | ✓ 1951ms | http |
| 183.239.109.98:22222 | 否 | ✓ 1469ms | ✓ 1113ms | ✓ 1431ms | ✓ 1112ms | http |

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
