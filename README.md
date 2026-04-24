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

最后更新：2026-04-24 21:48:11 UTC（2026-04-25 05:48:11 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1501ms | ✓ 1727ms | ✓ 1992ms | ✓ 1807ms | http |
| 152.70.91.193:40000 | ✓ 1789ms | 否 | 否 | ✓ 1868ms | ✓ 1990ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1887ms | ✓ 1658ms | ✓ 1622ms | ✓ 1298ms | http |
| 35.225.22.61:80 | ✓ 320ms | ✓ 1235ms | ✓ 406ms | ✓ 921ms | ✓ 885ms | http |
| 206.206.126.177:2412 | ✓ 910ms | ✓ 1651ms | ✓ 1222ms | ✓ 1125ms | ✓ 899ms | http |
| 120.92.108.86:7890 | ✓ 1995ms | 否 | ✓ 1662ms | ✓ 1838ms | ✓ 1879ms | http |
| 35.182.12.78:18278 | ✓ 755ms | 否 | ✓ 1637ms | 否 | ✓ 1501ms | http |
| 15.157.63.22:15433 | ✓ 667ms | 否 | ✓ 1674ms | 否 | ✓ 1651ms | http |
| 3.99.158.157:45219 | ✓ 998ms | 否 | ✓ 1530ms | 否 | ✓ 1719ms | http |
| 18.171.232.214:23578 | ✓ 685ms | 否 | ✓ 1660ms | 否 | ✓ 1339ms | http |
| 51.92.173.133:51177 | ✓ 1014ms | 否 | ✓ 944ms | 否 | ✓ 1779ms | http |
| 13.41.196.179:23714 | ✓ 1830ms | 否 | ✓ 790ms | ✓ 1883ms | 否 | http |
| 160.250.4.245:1 | ✓ 1743ms | 否 | ✓ 1610ms | ✓ 1359ms | ✓ 1192ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 761ms | ✓ 930ms | ✓ 1714ms | http |
| 46.101.95.183:8888 | ✓ 1846ms | 否 | ✓ 966ms | ✓ 1650ms | ✓ 1281ms | http |
| 177.93.132.244:3128 | ✓ 1569ms | 否 | ✓ 1823ms | 否 | ✓ 1755ms | http |
| 161.35.181.96:999 | ✓ 879ms | ✓ 1532ms | ✓ 539ms | ✓ 1186ms | ✓ 1018ms | http |
| 106.10.55.212:1121 | 否 | ✓ 1380ms | ✓ 1063ms | 否 | ✓ 1922ms | http |
| 103.55.225.34:8080 | ✓ 1464ms | 否 | ✓ 1394ms | ✓ 1865ms | ✓ 1538ms | http |
| 105.159.151.72:4699 | 否 | ✓ 1969ms | ✓ 1829ms | 否 | ✓ 1986ms | http |
| 45.153.231.229:8080 | ✓ 1428ms | ✓ 1842ms | ✓ 1642ms | ✓ 1739ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1396ms | ✓ 1663ms | ✓ 1383ms | 否 | http |
| 8.219.195.129:1080 | ✓ 1580ms | ✓ 1896ms | ✓ 959ms | ✓ 1231ms | 否 | http |
| 130.61.174.200:1080 | ✓ 1450ms | ✓ 1662ms | 否 | ✓ 1374ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1044ms | 否 | ✓ 1225ms | 否 | ✓ 1149ms | http |
| 121.230.8.17:1080 | ✓ 1079ms | ✓ 1472ms | ✓ 1242ms | ✓ 1583ms | ✓ 1252ms | http |
| 121.230.8.249:1080 | ✓ 1229ms | 否 | ✓ 1234ms | ✓ 1549ms | ✓ 1406ms | http |
| 2.27.54.161:1080 | ✓ 1359ms | 否 | ✓ 1899ms | ✓ 1843ms | ✓ 1472ms | http |
| 34.71.229.255:3128 | ✓ 1950ms | ✓ 1932ms | ✓ 1434ms | ✓ 1583ms | ✓ 1773ms | http |
| 45.186.6.104:3128 | ✓ 1853ms | ✓ 1783ms | ✓ 1763ms | 否 | 否 | http |
| 18.188.53.175:17723 | ✓ 1506ms | 否 | ✓ 817ms | 否 | ✓ 1649ms | http |
| 52.16.215.4:42899 | ✓ 749ms | 否 | ✓ 1489ms | ✓ 1994ms | ✓ 1469ms | http |
| 3.238.34.111:33720 | 否 | 否 | ✓ 800ms | ✓ 1991ms | ✓ 1740ms | http |
| 34.241.123.181:49254 | ✓ 983ms | 否 | ✓ 1226ms | ✓ 1965ms | ✓ 1693ms | http |
| 13.41.196.179:2130 | ✓ 1012ms | 否 | ✓ 1266ms | ✓ 1997ms | ✓ 1562ms | http |
| 34.246.223.187:4192 | ✓ 1044ms | 否 | ✓ 1170ms | ✓ 1840ms | ✓ 1916ms | http |
| 16.62.229.137:41511 | ✓ 1072ms | 否 | ✓ 1159ms | ✓ 1711ms | ✓ 1869ms | http |
| 40.176.49.172:5546 | ✓ 1340ms | 否 | ✓ 1398ms | 否 | ✓ 1976ms | http |
| 15.223.237.12:16965 | 否 | 否 | ✓ 1888ms | ✓ 1914ms | ✓ 1404ms | http |
| 3.19.213.118:40000 | ✓ 1818ms | 否 | ✓ 624ms | 否 | ✓ 1478ms | http |
| 3.14.146.121:17810 | ✓ 974ms | 否 | ✓ 1332ms | ✓ 1947ms | ✓ 1823ms | http |
| 3.137.167.45:93 | ✓ 809ms | 否 | ✓ 1691ms | 否 | ✓ 1598ms | http |
| 44.201.32.14:24330 | ✓ 1791ms | 否 | ✓ 822ms | ✓ 1988ms | ✓ 1453ms | http |
| 52.53.211.45:1111 | ✓ 1256ms | 否 | ✓ 1221ms | 否 | ✓ 1751ms | http |
| 54.229.201.146:48867 | ✓ 1189ms | 否 | ✓ 1394ms | 否 | ✓ 1389ms | http |
| 16.62.123.236:48789 | ✓ 851ms | 否 | ✓ 1605ms | ✓ 1835ms | ✓ 1614ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1832ms | ✓ 962ms | ✓ 1232ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1019ms | ✓ 1584ms | ✓ 1435ms | ✓ 1318ms | 否 | http |
| 183.232.248.73:7890 | ✓ 1028ms | ✓ 1266ms | ✓ 1017ms | ✓ 1246ms | ✓ 936ms | http |
| 167.71.196.178:80 | 否 | 否 | ✓ 1192ms | ✓ 1190ms | ✓ 960ms | http |
| 20.78.213.56:80 | 否 | ✓ 1797ms | ✓ 1614ms | ✓ 1101ms | ✓ 1055ms | http |
| 23.224.193.42:3128 | ✓ 557ms | 否 | ✓ 966ms | 否 | ✓ 1290ms | http |
| 23.224.193.43:3128 | ✓ 553ms | 否 | ✓ 1947ms | 否 | ✓ 696ms | http |
| 77.232.142.164:3128 | ✓ 1245ms | 否 | ✓ 960ms | 否 | ✓ 1460ms | http |
| 40.90.163.168:3128 | ✓ 1066ms | ✓ 1579ms | 否 | ✓ 1260ms | ✓ 1213ms | http |
| 23.224.193.46:3128 | ✓ 551ms | ✓ 1515ms | ✓ 1446ms | 否 | ✓ 1333ms | http |
| 58.63.109.230:10817 | ✓ 902ms | ✓ 1306ms | ✓ 1689ms | 否 | 否 | http |
| 8.209.238.110:47701 | ✓ 616ms | ✓ 1250ms | ✓ 827ms | ✓ 976ms | ✓ 766ms | http |
| 13.53.139.178:18161 | ✓ 1210ms | 否 | ✓ 1806ms | ✓ 1872ms | 否 | http |
| 35.180.75.159:11913 | 否 | 否 | ✓ 1723ms | ✓ 1719ms | ✓ 1423ms | http |
| 23.224.193.45:3128 | 否 | 否 | ✓ 1472ms | ✓ 1314ms | ✓ 1125ms | http |
| 106.44.155.90:2222 | ✓ 1546ms | ✓ 1573ms | 否 | 否 | ✓ 1427ms | http |
| 20.120.225.109:3128 | ✓ 777ms | ✓ 1505ms | ✓ 1694ms | ✓ 1149ms | ✓ 610ms | http |
| 210.223.44.230:3128 | ✓ 945ms | ✓ 1755ms | 否 | ✓ 1833ms | 否 | http |
| 43.132.188.134:443 | 否 | 否 | ✓ 1387ms | ✓ 1407ms | ✓ 977ms | http |
| 61.52.131.172:8443 | ✓ 985ms | 否 | ✓ 1063ms | 否 | ✓ 1121ms | http |
| 40.177.99.164:34583 | ✓ 1433ms | 否 | ✓ 1346ms | 否 | ✓ 1984ms | http |
| 13.60.163.108:38052 | ✓ 1192ms | 否 | ✓ 1888ms | ✓ 1595ms | 否 | http |
| 51.84.100.37:8118 | ✓ 1545ms | 否 | ✓ 1105ms | 否 | ✓ 1929ms | http |
| 103.39.51.207:8080 | ✓ 1653ms | 否 | 否 | ✓ 1610ms | ✓ 1624ms | http |

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
