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

最后更新：2026-05-06 19:12:30 UTC（2026-05-07 03:12:30 UTC+8）

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
| 218.108.131.186:17890 | ✓ 804ms | ✓ 1012ms | ✓ 809ms | ✓ 1048ms | ✓ 845ms | http |
| 113.160.132.26:8080 | ✓ 892ms | ✓ 1363ms | ✓ 903ms | ✓ 1237ms | ✓ 1287ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1103ms | ✓ 992ms | ✓ 1212ms | ✓ 904ms | http |
| 115.231.181.40:8128 | ✓ 819ms | ✓ 1036ms | ✓ 888ms | ✓ 1158ms | ✓ 928ms | http |
| 45.125.67.37:8443 | ✓ 822ms | 否 | ✓ 1048ms | ✓ 947ms | ✓ 1063ms | http |
| 120.92.212.16:7890 | ✓ 1035ms | 否 | ✓ 1036ms | ✓ 1441ms | ✓ 1193ms | http |
| 47.77.216.82:1080 | ✓ 650ms | 否 | ✓ 488ms | ✓ 1031ms | ✓ 1829ms | http |
| 45.153.231.229:8080 | ✓ 1266ms | ✓ 1929ms | ✓ 1941ms | 否 | ✓ 1732ms | http |
| 217.76.245.80:999 | ✓ 1333ms | ✓ 1840ms | ✓ 1348ms | 否 | ✓ 1436ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 493ms | ✓ 768ms | ✓ 586ms | http |
| 14.143.222.113:57788 | ✓ 932ms | 否 | ✓ 916ms | ✓ 1325ms | 否 | http |
| 185.191.236.162:3128 | ✓ 1521ms | 否 | ✓ 1917ms | 否 | ✓ 1925ms | http |
| 120.92.108.86:7890 | ✓ 1222ms | 否 | ✓ 1201ms | 否 | ✓ 1330ms | http |
| 122.224.198.218:808 | ✓ 1909ms | 否 | ✓ 1871ms | 否 | ✓ 1983ms | http |
| 20.78.26.206:8561 | 否 | 否 | ✓ 1471ms | ✓ 1696ms | ✓ 1756ms | http |
| 20.78.118.91:8561 | 否 | 否 | ✓ 1980ms | ✓ 1690ms | ✓ 1771ms | http |
| 20.210.39.153:8561 | 否 | 否 | ✓ 1986ms | ✓ 1689ms | ✓ 1774ms | http |
| 185.121.13.73:3128 | ✓ 1146ms | ✓ 1823ms | ✓ 1234ms | 否 | 否 | http |
| 34.96.238.40:8080 | ✓ 810ms | ✓ 1046ms | ✓ 1224ms | 否 | ✓ 1155ms | http |
| 43.156.132.113:3128 | ✓ 858ms | 否 | ✓ 694ms | ✓ 1008ms | ✓ 857ms | http |
| 190.12.150.244:999 | ✓ 1488ms | ✓ 1750ms | ✓ 1486ms | 否 | ✓ 1761ms | http |
| 152.70.91.193:40000 | 否 | ✓ 1839ms | ✓ 1424ms | ✓ 1325ms | ✓ 1328ms | http |
| 38.180.2.107:3128 | ✓ 1194ms | ✓ 1949ms | ✓ 1797ms | 否 | 否 | http |
| 46.39.105.157:8080 | ✓ 1038ms | 否 | ✓ 1132ms | 否 | ✓ 1716ms | http |
| 206.206.126.177:2412 | ✓ 1877ms | 否 | ✓ 1359ms | 否 | ✓ 1373ms | http |
| 20.27.13.35:8561 | 否 | ✓ 1488ms | ✓ 1353ms | ✓ 1766ms | 否 | http |
| 120.92.211.211:7890 | ✓ 966ms | 否 | ✓ 1951ms | 否 | ✓ 1122ms | http |
| 84.47.150.125:1080 | ✓ 1443ms | 否 | ✓ 1885ms | 否 | ✓ 1873ms | http |
| 43.133.44.89:8888 | ✓ 1505ms | 否 | ✓ 1961ms | 否 | ✓ 1738ms | http |
| 59.46.216.131:30001 | ✓ 1093ms | 否 | ✓ 1075ms | 否 | ✓ 1076ms | http |
| 137.59.47.73:3128 | ✓ 1525ms | 否 | 否 | ✓ 1692ms | ✓ 1505ms | http |
| 178.156.224.42:3128 | ✓ 1505ms | 否 | ✓ 1870ms | 否 | ✓ 1951ms | http |
| 62.113.119.14:8080 | ✓ 1297ms | ✓ 1729ms | ✓ 1205ms | 否 | 否 | http |
| 193.160.209.58:1080 | ✓ 1525ms | 否 | ✓ 1734ms | 否 | ✓ 1652ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1969ms | ✓ 1154ms | ✓ 1013ms | http |
| 158.160.215.167:8124 | ✓ 1705ms | 否 | ✓ 1320ms | 否 | ✓ 1956ms | http |
| 158.160.215.167:8127 | ✓ 1688ms | ✓ 1918ms | ✓ 1681ms | 否 | 否 | http |
| 80.92.204.47:1081 | ✓ 1719ms | 否 | ✓ 1363ms | 否 | ✓ 1599ms | http |
| 77.110.119.136:3128 | 否 | ✓ 1586ms | ✓ 437ms | ✓ 1672ms | 否 | http |
| 8.154.21.175:3128 | ✓ 818ms | ✓ 1055ms | ✓ 865ms | ✓ 1042ms | ✓ 867ms | http |
| 47.105.98.23:3128 | ✓ 999ms | ✓ 1093ms | ✓ 1840ms | 否 | ✓ 1050ms | http |
| 106.10.55.212:1121 | ✓ 1523ms | 否 | ✓ 1166ms | ✓ 1064ms | ✓ 1710ms | http |
| 82.114.228.67:1080 | ✓ 1642ms | ✓ 1866ms | 否 | 否 | ✓ 1360ms | http |
| 193.181.35.172:8118 | ✓ 1409ms | 否 | ✓ 921ms | 否 | ✓ 1987ms | http |
| 101.32.244.83:8080 | ✓ 976ms | 否 | ✓ 898ms | ✓ 1123ms | ✓ 1141ms | http |
| 121.43.196.210:8222 | ✓ 882ms | ✓ 979ms | ✓ 794ms | ✓ 1113ms | ✓ 893ms | http |
| 121.43.196.213:8222 | ✓ 925ms | ✓ 1010ms | ✓ 832ms | ✓ 1098ms | ✓ 909ms | http |
| 38.180.192.119:3128 | ✓ 686ms | ✓ 739ms | ✓ 785ms | ✓ 720ms | ✓ 760ms | http |
| 138.197.68.35:4857 | 否 | ✓ 1593ms | ✓ 731ms | 否 | ✓ 1738ms | http |
| 16.62.229.137:4853 | ✓ 831ms | 否 | ✓ 1167ms | ✓ 1796ms | 否 | http |
| 54.229.201.146:33577 | ✓ 876ms | 否 | ✓ 1751ms | 否 | ✓ 1889ms | http |
| 157.245.194.13:8888 | 否 | 否 | ✓ 976ms | ✓ 1003ms | ✓ 793ms | http |
| 195.26.243.76:3128 | ✓ 584ms | ✓ 1378ms | 否 | 否 | ✓ 1623ms | http |
| 91.233.223.147:3128 | ✓ 1085ms | 否 | ✓ 1683ms | 否 | ✓ 1750ms | http |
| 1.231.81.166:3128 | ✓ 1340ms | ✓ 1118ms | ✓ 1074ms | ✓ 1612ms | ✓ 1593ms | http |
| 38.194.254.134:999 | ✓ 1414ms | ✓ 1523ms | ✓ 1199ms | 否 | 否 | http |
| 14.143.222.113:10155 | ✓ 1581ms | 否 | ✓ 963ms | ✓ 1413ms | 否 | http |
| 3.101.133.120:80 | ✓ 488ms | ✓ 1319ms | ✓ 929ms | 否 | ✓ 838ms | http |
| 217.182.195.221:30000 | ✓ 1954ms | 否 | ✓ 1973ms | 否 | ✓ 1869ms | http |
| 117.236.124.166:3128 | ✓ 1667ms | 否 | ✓ 1854ms | 否 | ✓ 1389ms | http |
| 150.107.140.238:3128 | ✓ 1802ms | 否 | ✓ 783ms | 否 | ✓ 855ms | http |
| 80.92.204.47:1082 | ✓ 885ms | ✓ 1682ms | ✓ 851ms | 否 | 否 | http |
| 20.27.14.220:8561 | 否 | 否 | ✓ 582ms | ✓ 1319ms | ✓ 616ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 591ms | ✓ 1318ms | ✓ 618ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 590ms | ✓ 1320ms | ✓ 620ms | http |
| 158.160.215.167:8125 | ✓ 1342ms | ✓ 1933ms | ✓ 1543ms | 否 | 否 | http |
| 8.211.166.184:8081 | ✓ 1223ms | ✓ 1162ms | ✓ 731ms | ✓ 757ms | ✓ 610ms | http |
| 47.112.25.109:7890 | ✓ 1937ms | ✓ 1968ms | 否 | ✓ 1620ms | ✓ 1840ms | http |
| 77.110.107.80:1080 | 否 | ✓ 1936ms | ✓ 746ms | ✓ 1854ms | 否 | http |
| 20.18.193.135:8561 | ✓ 1578ms | ✓ 1354ms | ✓ 1277ms | ✓ 1781ms | 否 | http |
| 20.210.76.104:8561 | ✓ 1574ms | ✓ 1433ms | ✓ 1215ms | ✓ 1776ms | 否 | http |
| 20.210.76.178:8561 | ✓ 1578ms | ✓ 1362ms | ✓ 1280ms | ✓ 1790ms | 否 | http |
| 20.210.76.175:8561 | ✓ 1580ms | ✓ 1459ms | ✓ 1210ms | ✓ 1762ms | 否 | http |
| 20.27.15.49:8561 | ✓ 1258ms | ✓ 831ms | ✓ 705ms | ✓ 896ms | ✓ 711ms | http |
| 195.19.217.200:3128 | ✓ 1945ms | 否 | ✓ 1637ms | 否 | ✓ 1982ms | http |
| 116.171.106.26:3443 | ✓ 1343ms | ✓ 1391ms | 否 | ✓ 1566ms | ✓ 1332ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1219ms | ✓ 1435ms | ✓ 1749ms | 否 | http |
| 121.230.8.17:1080 | ✓ 981ms | ✓ 1357ms | ✓ 1009ms | ✓ 1273ms | ✓ 1181ms | http |
| 61.52.131.172:8443 | ✓ 900ms | ✓ 1208ms | ✓ 939ms | ✓ 1130ms | ✓ 904ms | http |
| 15.160.132.166:32605 | ✓ 1293ms | 否 | ✓ 1912ms | 否 | ✓ 1679ms | http |
| 86.104.74.110:1081 | ✓ 910ms | ✓ 1697ms | ✓ 1525ms | ✓ 1733ms | 否 | http |
| 116.171.106.15:3443 | 否 | ✓ 1360ms | ✓ 1401ms | ✓ 1430ms | 否 | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1432ms | ✓ 1636ms | ✓ 1610ms | http |
| 152.32.132.190:7890 | ✓ 1579ms | 否 | ✓ 1077ms | 否 | ✓ 1055ms | http |
| 103.125.16.12:8080 | ✓ 1894ms | ✓ 1966ms | ✓ 1977ms | ✓ 1363ms | 否 | http |

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
