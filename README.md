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

最后更新：2026-03-12 19:54:09 UTC（2026-03-13 03:54:09 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 761ms | ✓ 675ms | ✓ 891ms | ✓ 790ms | ✓ 646ms | http |
| 202.155.12.161:443 | ✓ 1147ms | 否 | ✓ 911ms | ✓ 1125ms | 否 | http |
| 205.209.118.30:3138 | ✓ 448ms | ✓ 1439ms | ✓ 1270ms | 否 | ✓ 1049ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1080ms | ✓ 1271ms | ✓ 923ms | ✓ 943ms | http |
| 113.160.132.26:8080 | ✓ 1318ms | ✓ 1456ms | ✓ 1263ms | ✓ 1654ms | ✓ 967ms | http |
| 193.168.173.136:443 | ✓ 1377ms | 否 | ✓ 1471ms | 否 | ✓ 1908ms | http |
| 171.251.172.78:5104 | ✓ 1312ms | 否 | ✓ 1441ms | ✓ 1558ms | ✓ 1625ms | http |
| 171.251.172.78:5106 | ✓ 1421ms | 否 | ✓ 1592ms | ✓ 1829ms | ✓ 1593ms | http |
| 116.80.96.104:3172 | ✓ 1491ms | 否 | 否 | ✓ 1838ms | ✓ 1646ms | http |
| 59.46.216.131:30001 | ✓ 1132ms | 否 | 否 | ✓ 1357ms | ✓ 1030ms | http |
| 120.92.212.16:7890 | ✓ 1618ms | 否 | ✓ 1915ms | ✓ 1392ms | 否 | http |
| 45.136.130.175:8443 | ✓ 368ms | ✓ 1011ms | ✓ 574ms | ✓ 655ms | ✓ 492ms | http |
| 45.136.131.47:8443 | ✓ 1094ms | ✓ 595ms | ✓ 743ms | ✓ 659ms | ✓ 524ms | http |
| 165.227.5.10:8888 | ✓ 367ms | ✓ 1674ms | ✓ 260ms | ✓ 877ms | ✓ 1799ms | http |
| 103.84.95.54:7890 | ✓ 685ms | 否 | ✓ 628ms | ✓ 821ms | ✓ 893ms | http |
| 35.225.22.61:80 | 否 | ✓ 1124ms | 否 | ✓ 1443ms | ✓ 1141ms | http |
| 81.70.169.194:80 | ✓ 951ms | ✓ 1187ms | ✓ 927ms | ✓ 1196ms | ✓ 993ms | http |
| 115.231.181.40:8128 | ✓ 930ms | ✓ 1067ms | ✓ 1989ms | ✓ 1091ms | ✓ 883ms | http |
| 190.9.109.198:999 | ✓ 679ms | ✓ 1590ms | ✓ 1296ms | ✓ 1770ms | ✓ 1317ms | http |
| 107.173.52.58:7890 | ✓ 1861ms | 否 | ✓ 1367ms | 否 | ✓ 1740ms | http |
| 138.124.53.25:7443 | ✓ 933ms | 否 | ✓ 1971ms | ✓ 1746ms | 否 | http |
| 202.141.161.53:10808 | ✓ 981ms | ✓ 1420ms | 否 | ✓ 1800ms | 否 | http |
| 101.43.255.96:80 | ✓ 933ms | ✓ 1188ms | 否 | ✓ 1238ms | ✓ 1012ms | http |
| 45.168.238.193:8443 | ✓ 359ms | ✓ 1898ms | ✓ 706ms | ✓ 1134ms | ✓ 902ms | http |
| 45.136.130.191:8443 | ✓ 469ms | ✓ 747ms | 否 | ✓ 1098ms | ✓ 698ms | http |
| 45.136.130.188:8443 | ✓ 519ms | ✓ 807ms | 否 | ✓ 1025ms | ✓ 679ms | http |
| 86.53.183.16:1080 | ✓ 1126ms | 否 | ✓ 1318ms | 否 | ✓ 1576ms | http |
| 178.236.245.17:3128 | ✓ 1181ms | 否 | ✓ 1048ms | ✓ 1878ms | ✓ 1412ms | http |
| 178.236.245.59:3128 | ✓ 1183ms | ✓ 1998ms | ✓ 1051ms | ✓ 1867ms | ✓ 1413ms | http |
| 91.107.141.42:8081 | ✓ 1382ms | 否 | ✓ 1426ms | 否 | ✓ 1942ms | http |
| 45.136.198.40:3128 | ✓ 1206ms | 否 | ✓ 1764ms | 否 | ✓ 1912ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1138ms | ✓ 868ms | 否 | ✓ 874ms | http |
| 45.136.130.223:8443 | ✓ 151ms | ✓ 738ms | ✓ 536ms | ✓ 688ms | ✓ 517ms | http |
| 62.113.119.14:8080 | ✓ 757ms | 否 | ✓ 823ms | ✓ 1667ms | ✓ 1306ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1118ms | ✓ 1132ms | ✓ 1858ms | http |
| 103.86.131.62:80 | ✓ 1001ms | 否 | 否 | ✓ 1317ms | ✓ 1602ms | http |
| 116.80.47.62:3172 | ✓ 1470ms | 否 | ✓ 1455ms | ✓ 1770ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1503ms | ✓ 1628ms | ✓ 1791ms | 否 | ✓ 963ms | http |
| 120.92.212.16:8890 | ✓ 940ms | 否 | ✓ 1806ms | 否 | ✓ 958ms | http |
| 34.101.184.164:3128 | ✓ 1021ms | 否 | ✓ 1534ms | ✓ 1474ms | ✓ 1015ms | http |
| 24.144.86.173:1080 | ✓ 953ms | ✓ 603ms | ✓ 266ms | ✓ 766ms | ✓ 596ms | http |
| 61.0.226.241:8080 | ✓ 1033ms | 否 | ✓ 1794ms | 否 | ✓ 1556ms | http |
| 8.219.97.248:80 | ✓ 1243ms | 否 | ✓ 1256ms | 否 | ✓ 1158ms | http |
| 152.42.213.210:8080 | ✓ 743ms | 否 | ✓ 1257ms | ✓ 1201ms | ✓ 864ms | http |
| 61.0.226.241:3128 | ✓ 1011ms | 否 | ✓ 985ms | 否 | ✓ 1549ms | http |
| 46.183.25.8:443 | ✓ 612ms | 否 | ✓ 277ms | ✓ 885ms | 否 | http |
| 14.225.212.37:7890 | ✓ 813ms | 否 | ✓ 1875ms | ✓ 1756ms | 否 | http |
| 116.80.49.168:3172 | ✓ 1487ms | 否 | ✓ 1458ms | ✓ 1925ms | 否 | http |
| 144.31.25.69:21064 | ✓ 1468ms | 否 | ✓ 1365ms | 否 | ✓ 1975ms | http |
| 103.150.101.15:1080 | ✓ 807ms | 否 | ✓ 1512ms | ✓ 1374ms | ✓ 905ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1111ms | 否 | ✓ 1671ms | ✓ 957ms | http |
| 45.140.147.155:1081 | ✓ 1662ms | ✓ 1279ms | ✓ 1331ms | 否 | ✓ 1257ms | http |
| 103.82.23.118:5247 | ✓ 1931ms | ✓ 1711ms | ✓ 1627ms | 否 | ✓ 1893ms | http |
| 103.39.51.190:8080 | ✓ 1848ms | 否 | 否 | ✓ 1741ms | ✓ 1535ms | http |
| 45.186.6.104:3128 | ✓ 1411ms | ✓ 1948ms | ✓ 1906ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1416ms | 否 | ✓ 1850ms | ✓ 1406ms | ✓ 1396ms | http |
| 165.225.113.220:11918 | ✓ 1116ms | 否 | ✓ 1142ms | ✓ 1650ms | ✓ 1067ms | http |
| 165.225.113.220:11180 | ✓ 1035ms | 否 | ✓ 966ms | ✓ 1214ms | ✓ 952ms | http |
| 165.225.113.220:10958 | ✓ 971ms | 否 | ✓ 1139ms | ✓ 1231ms | 否 | http |
| 103.166.255.170:8080 | ✓ 1763ms | 否 | ✓ 1602ms | ✓ 1582ms | ✓ 1414ms | http |
| 162.240.154.26:3128 | ✓ 594ms | ✓ 1508ms | 否 | ✓ 1862ms | 否 | http |
| 45.140.147.155:1082 | ✓ 995ms | 否 | ✓ 726ms | ✓ 1662ms | ✓ 1055ms | http |
| 190.52.104.214:999 | ✓ 1341ms | 否 | ✓ 1802ms | ✓ 1771ms | ✓ 1476ms | http |
| 152.42.213.210:443 | ✓ 1395ms | 否 | ✓ 1556ms | ✓ 1086ms | 否 | http |
| 137.184.6.117:3128 | ✓ 1515ms | ✓ 1880ms | ✓ 1354ms | ✓ 636ms | ✓ 484ms | http |
| 45.136.130.239:8443 | 否 | ✓ 740ms | 否 | ✓ 674ms | ✓ 517ms | http |
| 121.230.8.97:1080 | ✓ 1012ms | ✓ 1341ms | 否 | ✓ 1416ms | ✓ 926ms | http |
| 137.184.1.87:3128 | 否 | 否 | ✓ 1603ms | ✓ 1360ms | ✓ 491ms | http |
| 8.222.175.80:6128 | ✓ 690ms | ✓ 1758ms | ✓ 687ms | ✓ 1014ms | ✓ 801ms | http |
| 121.230.9.19:1080 | ✓ 1273ms | ✓ 1247ms | ✓ 1519ms | ✓ 1402ms | ✓ 1022ms | http |
| 106.14.203.63:3333 | ✓ 782ms | ✓ 1591ms | 否 | 否 | ✓ 1036ms | http |
| 146.56.182.165:3128 | ✓ 889ms | 否 | 否 | ✓ 1674ms | ✓ 1394ms | http |
| 180.103.19.219:1080 | ✓ 1251ms | ✓ 1562ms | ✓ 1425ms | ✓ 1630ms | ✓ 1162ms | http |
| 47.105.98.23:3128 | ✓ 1117ms | ✓ 1219ms | ✓ 1995ms | ✓ 1216ms | ✓ 1705ms | http |
| 222.109.119.178:3128 | ✓ 693ms | ✓ 1340ms | ✓ 1376ms | 否 | 否 | http |
| 187.102.195.53:999 | ✓ 1134ms | 否 | ✓ 853ms | 否 | ✓ 1792ms | http |
| 103.172.0.78:8080 | 否 | 否 | ✓ 1380ms | ✓ 1545ms | ✓ 1591ms | http |
| 116.63.160.98:8899 | ✓ 966ms | ✓ 1152ms | ✓ 921ms | ✓ 1190ms | ✓ 1185ms | http |
| 47.77.193.180:1080 | ✓ 723ms | ✓ 731ms | ✓ 199ms | ✓ 756ms | ✓ 493ms | http |
| 202.129.206.239:3128 | ✓ 1002ms | 否 | ✓ 1823ms | ✓ 1423ms | ✓ 1286ms | http |
| 120.55.163.237:10086 | ✓ 784ms | ✓ 982ms | ✓ 881ms | ✓ 1017ms | ✓ 1059ms | http |
| 116.80.49.156:3172 | ✓ 1780ms | 否 | ✓ 1882ms | ✓ 1830ms | 否 | http |
| 110.172.29.131:3128 | ✓ 1774ms | ✓ 1981ms | ✓ 1508ms | ✓ 1178ms | ✓ 929ms | http |
| 116.80.96.106:3172 | ✓ 1595ms | 否 | ✓ 1913ms | 否 | ✓ 1790ms | http |
| 116.80.49.169:3172 | ✓ 1594ms | 否 | ✓ 1940ms | ✓ 1990ms | 否 | http |
| 116.80.49.162:3172 | ✓ 1490ms | 否 | ✓ 1439ms | ✓ 1822ms | 否 | http |
| 112.163.160.93:3128 | 否 | ✓ 1486ms | ✓ 1013ms | ✓ 983ms | ✓ 746ms | http |
| 111.79.111.126:3128 | ✓ 1636ms | 否 | 否 | ✓ 1862ms | ✓ 1201ms | http |

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
