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

最后更新：2026-04-17 11:54:29 UTC（2026-04-17 19:54:29 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:8080 | ✓ 1119ms | ✓ 1902ms | 否 | ✓ 1227ms | 否 | http |
| 188.246.224.49:7890 | ✓ 799ms | 否 | ✓ 1052ms | ✓ 1632ms | 否 | http |
| 45.140.147.82:1081 | ✓ 683ms | 否 | ✓ 948ms | ✓ 1556ms | ✓ 1065ms | http |
| 185.138.116.150:8080 | 否 | 否 | ✓ 1659ms | ✓ 1434ms | ✓ 1019ms | http |
| 113.160.132.26:8080 | ✓ 1719ms | ✓ 1848ms | 否 | ✓ 1410ms | ✓ 1183ms | http |
| 34.71.229.255:3128 | ✓ 1749ms | 否 | ✓ 1684ms | ✓ 1533ms | ✓ 1206ms | http |
| 212.58.132.5:8888 | ✓ 1563ms | 否 | ✓ 1572ms | ✓ 1446ms | ✓ 1161ms | http |
| 168.144.75.9:3128 | ✓ 1586ms | 否 | ✓ 1831ms | ✓ 1841ms | ✓ 1150ms | http |
| 1.231.81.166:3128 | ✓ 1999ms | 否 | ✓ 1915ms | ✓ 1711ms | ✓ 1440ms | http |
| 149.104.4.88:10809 | ✓ 1903ms | 否 | ✓ 1805ms | ✓ 1499ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1175ms | ✓ 1630ms | 否 | 否 | ✓ 1972ms | http |
| 149.51.42.10:3128 | ✓ 943ms | ✓ 1229ms | 否 | ✓ 1743ms | 否 | http |
| 130.61.30.221:8080 | ✓ 666ms | 否 | ✓ 1129ms | ✓ 1812ms | 否 | http |
| 177.93.132.244:3128 | ✓ 664ms | 否 | ✓ 631ms | ✓ 1906ms | ✓ 1452ms | http |
| 218.108.131.186:17890 | ✓ 1029ms | ✓ 1297ms | ✓ 1026ms | ✓ 1354ms | ✓ 1074ms | http |
| 120.92.108.86:7890 | ✓ 1509ms | 否 | ✓ 1399ms | ✓ 1920ms | ✓ 1852ms | http |
| 157.230.178.216:8088 | ✓ 514ms | 否 | ✓ 1023ms | ✓ 1375ms | ✓ 1372ms | http |
| 117.236.124.166:3128 | ✓ 1384ms | 否 | ✓ 1931ms | 否 | ✓ 1752ms | http |
| 158.160.215.167:8124 | ✓ 792ms | 否 | ✓ 984ms | 否 | ✓ 1961ms | http |
| 84.47.150.125:1080 | ✓ 980ms | 否 | ✓ 1315ms | 否 | ✓ 1741ms | http |
| 103.85.113.66:9999 | ✓ 1661ms | ✓ 1635ms | ✓ 1544ms | 否 | ✓ 1418ms | http |
| 84.47.150.126:1080 | ✓ 1126ms | 否 | ✓ 1730ms | 否 | ✓ 1854ms | http |
| 45.140.147.155:1082 | ✓ 568ms | 否 | ✓ 427ms | ✓ 1534ms | ✓ 1101ms | http |
| 20.127.128.70:8080 | ✓ 542ms | ✓ 1923ms | ✓ 1067ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 556ms | ✓ 1782ms | 否 | ✓ 1553ms | 否 | http |
| 43.132.188.134:443 | ✓ 840ms | ✓ 1147ms | 否 | 否 | ✓ 842ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1454ms | ✓ 1497ms | 否 | ✓ 1902ms | http |
| 78.11.96.22:8888 | ✓ 959ms | ✓ 1617ms | ✓ 1193ms | ✓ 1793ms | ✓ 1390ms | http |
| 103.113.70.189:1082 | ✓ 863ms | 否 | ✓ 809ms | ✓ 957ms | ✓ 878ms | http |
| 15.161.131.175:3129 | ✓ 659ms | 否 | ✓ 1346ms | 否 | ✓ 1756ms | http |
| 144.31.27.49:1080 | ✓ 1563ms | 否 | ✓ 1505ms | 否 | ✓ 1574ms | http |
| 45.12.151.226:2829 | ✓ 1255ms | 否 | ✓ 1943ms | 否 | ✓ 1825ms | http |
| 35.225.22.61:80 | ✓ 864ms | 否 | ✓ 952ms | ✓ 1168ms | ✓ 1039ms | http |
| 147.45.214.210:1080 | ✓ 738ms | 否 | ✓ 534ms | 否 | ✓ 1240ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 830ms | ✓ 1170ms | ✓ 928ms | http |
| 108.131.109.106:44593 | 否 | 否 | ✓ 1130ms | ✓ 1794ms | ✓ 1456ms | http |
| 166.249.54.61:7234 | 否 | ✓ 1821ms | 否 | ✓ 1757ms | ✓ 1733ms | http |
| 185.114.73.2:1080 | ✓ 1177ms | ✓ 1816ms | ✓ 1263ms | 否 | 否 | http |
| 45.140.147.82:1082 | ✓ 542ms | 否 | ✓ 868ms | ✓ 1191ms | 否 | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 205ms | ✓ 1156ms | ✓ 938ms | http |
| 34.96.238.40:8080 | ✓ 1066ms | 否 | ✓ 1417ms | ✓ 1232ms | 否 | http |
| 104.168.93.120:8080 | ✓ 1152ms | 否 | ✓ 1327ms | ✓ 1693ms | 否 | http |
| 101.32.243.189:80 | 否 | ✓ 1950ms | ✓ 1849ms | ✓ 1678ms | ✓ 1552ms | http |
| 34.246.183.20:3128 | ✓ 831ms | 否 | ✓ 1709ms | ✓ 1813ms | ✓ 1539ms | http |
| 34.246.223.187:3128 | ✓ 1588ms | 否 | ✓ 1284ms | ✓ 1920ms | 否 | http |
| 202.141.161.53:10808 | 否 | ✓ 1407ms | ✓ 1997ms | ✓ 1507ms | 否 | http |
| 217.52.247.89:1976 | 否 | 否 | ✓ 1930ms | ✓ 1923ms | ✓ 1673ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1522ms | ✓ 1825ms | ✓ 1444ms | ✓ 1446ms | http |
| 104.247.51.76:3128 | ✓ 413ms | 否 | ✓ 1181ms | ✓ 1037ms | ✓ 763ms | http |
| 131.196.245.120:999 | 否 | 否 | ✓ 460ms | ✓ 1820ms | ✓ 1645ms | http |
| 59.46.216.131:30001 | ✓ 1333ms | ✓ 1680ms | ✓ 1300ms | 否 | 否 | http |
| 15.161.131.175:19114 | 否 | 否 | ✓ 1697ms | ✓ 1620ms | ✓ 1797ms | http |
| 51.95.13.205:12 | ✓ 886ms | 否 | ✓ 1674ms | 否 | ✓ 1961ms | http |
| 157.0.142.246:10061 | ✓ 1254ms | ✓ 1493ms | ✓ 1176ms | ✓ 1562ms | ✓ 1332ms | http |
| 42.101.8.101:8888 | ✓ 1385ms | ✓ 1668ms | ✓ 1522ms | ✓ 1728ms | 否 | http |
| 158.160.215.167:8127 | ✓ 1223ms | 否 | ✓ 1434ms | 否 | ✓ 1223ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 556ms | ✓ 1472ms | ✓ 1050ms | http |
| 159.223.225.118:8888 | ✓ 957ms | ✓ 1508ms | ✓ 632ms | 否 | 否 | http |
| 2.27.18.184:1080 | ✓ 1327ms | ✓ 1656ms | 否 | 否 | ✓ 1695ms | http |
| 212.34.146.118:3128 | 否 | 否 | ✓ 1329ms | ✓ 1958ms | ✓ 1190ms | http |
| 103.138.70.165:3129 | 否 | 否 | ✓ 1829ms | ✓ 1903ms | ✓ 1886ms | http |
| 133.18.123.225:26021 | ✓ 1814ms | 否 | ✓ 1414ms | ✓ 1621ms | ✓ 1430ms | http |
| 116.58.161.203:26021 | ✓ 1820ms | 否 | ✓ 1410ms | ✓ 1583ms | ✓ 1540ms | http |
| 94.131.118.129:1081 | ✓ 733ms | 否 | ✓ 779ms | ✓ 1469ms | ✓ 1204ms | http |
| 94.242.50.157:8888 | ✓ 1295ms | ✓ 1790ms | ✓ 1681ms | 否 | 否 | http |
| 167.71.196.178:80 | 否 | 否 | ✓ 922ms | ✓ 1289ms | ✓ 1041ms | http |
| 217.77.102.18:3128 | 否 | ✓ 1790ms | ✓ 1006ms | 否 | ✓ 1698ms | http |
| 47.101.159.19:8899 | ✓ 1073ms | 否 | 否 | ✓ 1349ms | ✓ 1075ms | http |
| 110.172.29.131:3128 | 否 | 否 | ✓ 1600ms | ✓ 1404ms | ✓ 1121ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1766ms | ✓ 1799ms | ✓ 1998ms | 否 | http |
| 62.234.206.73:3128 | 否 | ✓ 1425ms | ✓ 1198ms | 否 | ✓ 1272ms | http |
| 101.255.211.213:8080 | 否 | 否 | ✓ 1487ms | ✓ 1739ms | ✓ 1690ms | http |
| 5.253.43.103:3128 | ✓ 1823ms | 否 | ✓ 1241ms | 否 | ✓ 1366ms | http |
| 138.124.99.216:8888 | ✓ 1005ms | 否 | 否 | ✓ 1975ms | ✓ 1287ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1365ms | ✓ 1757ms | ✓ 1826ms | ✓ 1117ms | http |
| 94.131.118.39:1081 | ✓ 1921ms | 否 | 否 | ✓ 1914ms | ✓ 1868ms | http |
| 217.52.247.89:1981 | 否 | 否 | ✓ 1848ms | ✓ 1823ms | ✓ 1912ms | http |
| 8.219.97.248:80 | ✓ 1387ms | 否 | ✓ 1535ms | ✓ 1644ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1035ms | ✓ 1407ms | ✓ 1175ms | ✓ 1434ms | ✓ 1131ms | http |
| 108.131.109.106:1516 | 否 | 否 | ✓ 1277ms | ✓ 1715ms | ✓ 1680ms | http |
| 103.166.158.99:8080 | 否 | 否 | ✓ 1515ms | ✓ 1886ms | ✓ 1704ms | http |
| 160.250.5.22:1 | ✓ 1788ms | 否 | ✓ 1535ms | ✓ 1468ms | ✓ 1156ms | http |

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
