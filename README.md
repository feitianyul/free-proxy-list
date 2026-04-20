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

最后更新：2026-04-20 07:12:03 UTC（2026-04-20 15:12:03 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:3128 | ✓ 1352ms | ✓ 1463ms | 否 | ✓ 1666ms | 否 | http |
| 47.84.73.61:1080 | ✓ 1144ms | 否 | ✓ 900ms | ✓ 1044ms | ✓ 811ms | http |
| 45.76.207.177:40000 | ✓ 1467ms | 否 | ✓ 568ms | 否 | ✓ 770ms | http |
| 113.160.132.26:8080 | ✓ 1710ms | ✓ 1371ms | ✓ 1279ms | ✓ 1249ms | ✓ 1055ms | http |
| 115.231.181.40:8128 | ✓ 974ms | ✓ 1065ms | ✓ 982ms | 否 | 否 | http |
| 46.101.95.183:8888 | ✓ 1585ms | 否 | ✓ 949ms | ✓ 1764ms | ✓ 1409ms | http |
| 152.42.208.139:8118 | ✓ 1145ms | 否 | ✓ 1305ms | ✓ 1068ms | ✓ 807ms | http |
| 152.70.91.193:40000 | ✓ 1576ms | 否 | 否 | ✓ 1773ms | ✓ 1757ms | http |
| 43.132.188.134:443 | ✓ 1463ms | 否 | ✓ 1668ms | ✓ 826ms | 否 | http |
| 35.225.22.61:80 | ✓ 376ms | ✓ 1698ms | 否 | 否 | ✓ 934ms | http |
| 149.51.42.10:8080 | ✓ 585ms | ✓ 1476ms | 否 | ✓ 1501ms | 否 | http |
| 34.96.238.40:8080 | ✓ 824ms | ✓ 1046ms | 否 | ✓ 1099ms | ✓ 1027ms | http |
| 192.3.248.190:8014 | ✓ 491ms | 否 | ✓ 1247ms | ✓ 1424ms | ✓ 827ms | http |
| 218.108.131.186:17890 | ✓ 821ms | ✓ 982ms | ✓ 866ms | ✓ 1053ms | ✓ 852ms | http |
| 168.110.52.228:3128 | ✓ 1012ms | 否 | ✓ 1856ms | ✓ 956ms | ✓ 777ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1147ms | ✓ 939ms | 否 | ✓ 946ms | http |
| 36.141.21.200:7890 | ✓ 890ms | ✓ 1135ms | ✓ 948ms | ✓ 1872ms | ✓ 973ms | http |
| 223.84.151.86:30005 | ✓ 1554ms | ✓ 1158ms | ✓ 998ms | ✓ 1398ms | ✓ 1177ms | http |
| 194.104.9.38:3128 | ✓ 705ms | ✓ 1791ms | ✓ 1593ms | 否 | ✓ 1758ms | http |
| 188.246.224.49:7890 | ✓ 887ms | 否 | ✓ 1886ms | 否 | ✓ 1221ms | http |
| 91.99.15.45:2095 | ✓ 671ms | ✓ 1956ms | ✓ 1717ms | 否 | ✓ 1947ms | http |
| 208.87.243.199:7878 | ✓ 914ms | 否 | ✓ 1143ms | 否 | ✓ 1536ms | http |
| 113.176.92.71:3128 | ✓ 1097ms | ✓ 1298ms | ✓ 1226ms | ✓ 1405ms | ✓ 942ms | http |
| 185.138.116.150:8080 | ✓ 1087ms | 否 | 否 | ✓ 1817ms | ✓ 1345ms | http |
| 45.153.231.229:8080 | 否 | ✓ 1965ms | ✓ 1480ms | 否 | ✓ 1910ms | http |
| 168.144.75.9:3128 | ✓ 1281ms | 否 | ✓ 1402ms | 否 | ✓ 1923ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1386ms | ✓ 1099ms | 否 | ✓ 1101ms | http |
| 213.32.85.26:3128 | ✓ 700ms | ✓ 1948ms | ✓ 1386ms | 否 | 否 | http |
| 177.93.132.244:3128 | ✓ 1558ms | 否 | ✓ 833ms | 否 | ✓ 1846ms | http |
| 159.223.225.118:8888 | ✓ 1215ms | 否 | ✓ 1652ms | ✓ 1836ms | ✓ 1901ms | http |
| 84.47.150.125:1080 | ✓ 1251ms | 否 | ✓ 1346ms | 否 | ✓ 1805ms | http |
| 195.26.224.49:3128 | ✓ 1224ms | 否 | ✓ 815ms | ✓ 1708ms | ✓ 1918ms | http |
| 120.92.212.16:8890 | ✓ 920ms | 否 | ✓ 845ms | ✓ 1174ms | ✓ 1004ms | http |
| 81.30.156.115:8080 | ✓ 795ms | ✓ 1874ms | ✓ 1495ms | 否 | ✓ 1665ms | http |
| 120.92.108.86:7890 | ✓ 1181ms | 否 | ✓ 1180ms | ✓ 1886ms | ✓ 1667ms | http |
| 107.150.41.226:18080 | ✓ 903ms | ✓ 1464ms | ✓ 489ms | ✓ 1093ms | ✓ 925ms | http |
| 52.221.190.128:1080 | ✓ 1604ms | ✓ 1094ms | ✓ 952ms | ✓ 1020ms | ✓ 872ms | http |
| 1.231.81.166:3128 | ✓ 1626ms | 否 | ✓ 1172ms | ✓ 877ms | ✓ 738ms | http |
| 212.58.132.5:8888 | ✓ 1596ms | 否 | ✓ 1483ms | ✓ 1827ms | ✓ 1211ms | http |
| 42.101.8.101:8888 | ✓ 1192ms | ✓ 1511ms | ✓ 1310ms | ✓ 1430ms | ✓ 1189ms | http |
| 150.249.255.91:3128 | 否 | ✓ 893ms | 否 | ✓ 794ms | ✓ 604ms | http |
| 84.47.150.126:1080 | ✓ 856ms | 否 | ✓ 1600ms | 否 | ✓ 1751ms | http |
| 111.79.111.126:3128 | ✓ 1204ms | ✓ 1745ms | ✓ 902ms | ✓ 1093ms | ✓ 1041ms | http |
| 144.31.27.49:1080 | ✓ 1192ms | 否 | ✓ 1656ms | 否 | ✓ 1903ms | http |
| 183.98.143.134:8039 | ✓ 684ms | ✓ 1741ms | ✓ 1018ms | ✓ 987ms | ✓ 1623ms | http |
| 8.219.97.248:80 | ✓ 1108ms | 否 | ✓ 899ms | ✓ 1490ms | ✓ 1197ms | http |
| 144.31.25.69:21064 | ✓ 1416ms | 否 | ✓ 1451ms | 否 | ✓ 1997ms | http |
| 62.113.119.14:8080 | ✓ 775ms | 否 | ✓ 771ms | ✓ 1654ms | ✓ 1263ms | http |
| 102.134.49.165:6005 | ✓ 1336ms | ✓ 1197ms | ✓ 1890ms | ✓ 1079ms | ✓ 717ms | http |
| 102.134.48.240:6005 | ✓ 1336ms | ✓ 1215ms | 否 | ✓ 1262ms | ✓ 941ms | http |
| 121.230.9.248:1080 | ✓ 1361ms | ✓ 1338ms | ✓ 934ms | ✓ 1376ms | ✓ 1661ms | http |
| 121.230.8.227:1080 | ✓ 1522ms | ✓ 1480ms | ✓ 1729ms | 否 | ✓ 1578ms | http |
| 57.128.188.167:9172 | ✓ 1875ms | 否 | ✓ 1656ms | 否 | ✓ 1757ms | http |
| 78.11.96.22:8888 | 否 | 否 | ✓ 1499ms | ✓ 1721ms | ✓ 1563ms | http |
| 121.232.73.214:1080 | ✓ 1122ms | ✓ 1103ms | ✓ 977ms | ✓ 1485ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1391ms | 否 | ✓ 1035ms | 否 | ✓ 1873ms | http |
| 94.131.118.39:1081 | ✓ 1942ms | 否 | ✓ 1911ms | 否 | ✓ 1614ms | http |
| 159.89.191.221:3128 | ✓ 1434ms | 否 | ✓ 1317ms | ✓ 1329ms | ✓ 969ms | http |
| 20.127.128.70:8080 | ✓ 1522ms | 否 | ✓ 1834ms | 否 | ✓ 1787ms | http |
| 162.240.154.26:3128 | ✓ 901ms | ✓ 1704ms | ✓ 1539ms | ✓ 1358ms | ✓ 1165ms | http |
| 94.131.118.39:1082 | ✓ 1034ms | 否 | ✓ 1073ms | 否 | ✓ 1806ms | http |
| 220.197.44.36:3128 | 否 | ✓ 1704ms | ✓ 1779ms | ✓ 1748ms | 否 | http |
| 116.171.106.15:3443 | 否 | 否 | ✓ 1856ms | ✓ 1565ms | ✓ 1475ms | http |
| 202.141.161.53:10808 | ✓ 1848ms | 否 | ✓ 1492ms | ✓ 1281ms | ✓ 1011ms | http |
| 172.105.118.164:3128 | ✓ 1398ms | 否 | ✓ 1419ms | ✓ 1377ms | ✓ 818ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1293ms | ✓ 1930ms | ✓ 1967ms | ✓ 1277ms | http |
| 117.122.240.82:3338 | 否 | ✓ 1475ms | ✓ 1049ms | ✓ 1362ms | ✓ 1234ms | http |
| 85.190.99.143:443 | ✓ 1301ms | 否 | ✓ 1595ms | 否 | ✓ 1862ms | http |
| 47.112.25.109:7890 | ✓ 1898ms | 否 | 否 | ✓ 1643ms | ✓ 1971ms | http |
| 203.150.128.61:8080 | ✓ 1757ms | 否 | ✓ 1594ms | 否 | ✓ 1788ms | http |
| 23.138.80.49:999 | 否 | 否 | ✓ 1448ms | ✓ 1928ms | ✓ 1980ms | http |
| 47.85.84.192:20000 | ✓ 1649ms | 否 | 否 | ✓ 1262ms | ✓ 1602ms | http |
| 202.129.206.239:3128 | ✓ 1578ms | 否 | ✓ 1343ms | ✓ 1459ms | ✓ 1387ms | http |
| 202.40.185.146:2025 | ✓ 1727ms | 否 | ✓ 1809ms | ✓ 1881ms | ✓ 1574ms | http |
| 117.236.124.166:3128 | ✓ 1925ms | 否 | ✓ 1877ms | 否 | ✓ 1747ms | http |
| 120.92.211.211:7890 | ✓ 1532ms | 否 | ✓ 1934ms | 否 | ✓ 1827ms | http |
| 116.171.106.26:3443 | 否 | ✓ 1881ms | ✓ 1752ms | 否 | ✓ 1992ms | http |
| 116.171.106.78:3443 | ✓ 1559ms | ✓ 1576ms | 否 | ✓ 1815ms | ✓ 1632ms | http |
| 103.133.254.4:3128 | ✓ 1034ms | 否 | ✓ 1818ms | ✓ 1611ms | ✓ 1185ms | http |
| 61.52.131.172:8443 | ✓ 825ms | ✓ 1101ms | ✓ 938ms | ✓ 1158ms | ✓ 916ms | http |
| 8.140.104.98:3128 | ✓ 845ms | ✓ 1136ms | ✓ 931ms | ✓ 1188ms | ✓ 1054ms | http |
| 101.32.243.189:80 | 否 | 否 | ✓ 1436ms | ✓ 1330ms | ✓ 1190ms | http |
| 103.39.51.207:8080 | ✓ 1359ms | 否 | 否 | ✓ 1833ms | ✓ 1552ms | http |
| 160.250.5.22:1 | ✓ 1538ms | 否 | ✓ 1384ms | ✓ 1513ms | ✓ 923ms | http |

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
