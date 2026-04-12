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

最后更新：2026-04-12 12:33:26 UTC（2026-04-12 20:33:26 UTC+8）

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
| 5.196.101.18:3128 | ✓ 592ms | ✓ 1560ms | ✓ 875ms | ✓ 1864ms | ✓ 1520ms | http |
| 147.161.210.140:8800 | ✓ 1884ms | 否 | ✓ 739ms | ✓ 1477ms | ✓ 1014ms | http |
| 167.103.115.102:8800 | ✓ 1685ms | 否 | ✓ 1224ms | ✓ 1278ms | ✓ 1144ms | http |
| 113.160.132.26:8080 | ✓ 1943ms | 否 | ✓ 962ms | ✓ 1864ms | ✓ 1441ms | http |
| 159.223.225.118:8888 | ✓ 1186ms | 否 | ✓ 1332ms | 否 | ✓ 1898ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1769ms | ✓ 1538ms | ✓ 1998ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1395ms | ✓ 1574ms | ✓ 1267ms | http |
| 167.103.34.108:8800 | ✓ 1590ms | 否 | ✓ 1616ms | ✓ 1687ms | ✓ 1422ms | http |
| 167.103.144.127:8800 | ✓ 1362ms | 否 | ✓ 1618ms | ✓ 1844ms | 否 | http |
| 137.59.47.73:3128 | ✓ 871ms | 否 | ✓ 1066ms | ✓ 1601ms | ✓ 951ms | http |
| 46.30.46.133:3128 | ✓ 514ms | 否 | ✓ 1486ms | ✓ 1483ms | 否 | http |
| 45.167.125.21:999 | ✓ 1979ms | 否 | ✓ 1271ms | 否 | ✓ 1752ms | http |
| 162.240.154.26:3128 | ✓ 492ms | ✓ 986ms | ✓ 412ms | ✓ 918ms | ✓ 689ms | http |
| 210.223.44.230:3128 | ✓ 1582ms | ✓ 1050ms | ✓ 746ms | ✓ 1045ms | ✓ 757ms | http |
| 103.82.23.118:5216 | ✓ 1621ms | 否 | ✓ 1457ms | ✓ 1783ms | ✓ 1576ms | http |
| 167.103.31.122:8800 | ✓ 1730ms | 否 | ✓ 1458ms | 否 | ✓ 1677ms | http |
| 5.104.87.17:8051 | ✓ 1737ms | 否 | 否 | ✓ 1773ms | ✓ 1342ms | http |
| 147.161.239.240:8800 | ✓ 601ms | ✓ 1815ms | ✓ 1046ms | ✓ 1539ms | ✓ 1535ms | http |
| 94.72.109.214:8888 | ✓ 764ms | 否 | ✓ 770ms | ✓ 1997ms | ✓ 1143ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1792ms | ✓ 1924ms | ✓ 1298ms | http |
| 95.214.9.93:3128 | ✓ 977ms | 否 | 否 | ✓ 1952ms | ✓ 1491ms | http |
| 120.92.108.86:7890 | ✓ 1367ms | 否 | ✓ 1921ms | 否 | ✓ 1565ms | http |
| 45.136.131.54:8448 | ✓ 1251ms | 否 | ✓ 1547ms | ✓ 843ms | 否 | http |
| 38.145.208.226:8448 | ✓ 1513ms | ✓ 941ms | 否 | 否 | ✓ 1963ms | http |
| 45.136.130.181:8445 | 否 | ✓ 950ms | ✓ 1252ms | 否 | ✓ 645ms | http |
| 202.141.161.53:10808 | ✓ 1231ms | ✓ 1551ms | 否 | ✓ 1590ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1784ms | 否 | 否 | ✓ 1644ms | ✓ 1320ms | http |
| 1.231.81.166:3128 | ✓ 1567ms | ✓ 1639ms | ✓ 1728ms | ✓ 1628ms | 否 | http |
| 185.76.241.110:10001 | ✓ 1318ms | 否 | ✓ 1686ms | 否 | ✓ 1898ms | http |
| 59.46.216.131:30001 | ✓ 1135ms | ✓ 1917ms | 否 | 否 | ✓ 1174ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 569ms | ✓ 1124ms | ✓ 1115ms | http |
| 38.34.179.174:8453 | ✓ 1280ms | 否 | ✓ 1266ms | ✓ 992ms | ✓ 1460ms | http |
| 194.67.99.223:1080 | ✓ 617ms | ✓ 1909ms | 否 | 否 | ✓ 1896ms | http |
| 37.187.109.70:10111 | ✓ 669ms | ✓ 1700ms | ✓ 1206ms | 否 | 否 | http |
| 223.84.151.86:30005 | ✓ 1583ms | 否 | ✓ 1133ms | ✓ 1466ms | ✓ 1597ms | http |
| 36.103.198.235:7890 | 否 | ✓ 1429ms | ✓ 1867ms | ✓ 1375ms | ✓ 1211ms | http |
| 8.219.97.248:80 | ✓ 1124ms | 否 | ✓ 1242ms | ✓ 1559ms | ✓ 1684ms | http |
| 36.141.21.200:7890 | ✓ 1205ms | ✓ 1365ms | 否 | ✓ 1600ms | 否 | http |
| 38.34.179.229:8447 | 否 | ✓ 1105ms | 否 | ✓ 1681ms | ✓ 1427ms | http |
| 38.145.203.111:8453 | 否 | ✓ 1600ms | ✓ 1866ms | ✓ 1951ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1501ms | 否 | ✓ 1273ms | 否 | ✓ 1095ms | http |
| 38.145.203.124:8445 | 否 | ✓ 1752ms | ✓ 1719ms | ✓ 1141ms | 否 | http |
| 45.149.92.147:5001 | ✓ 725ms | 否 | ✓ 1579ms | ✓ 1500ms | ✓ 766ms | http |
| 109.107.179.140:31000 | ✓ 1216ms | 否 | ✓ 1283ms | ✓ 1677ms | ✓ 1961ms | http |
| 107.172.102.234:40621 | ✓ 800ms | ✓ 1545ms | ✓ 824ms | ✓ 892ms | ✓ 791ms | http |
| 34.85.118.216:3128 | ✓ 1342ms | ✓ 1781ms | ✓ 617ms | ✓ 946ms | ✓ 990ms | http |
| 54.222.174.194:80 | ✓ 1630ms | ✓ 1924ms | ✓ 1611ms | 否 | 否 | http |
| 20.78.213.56:80 | ✓ 1952ms | 否 | 否 | ✓ 1592ms | ✓ 1901ms | http |
| 136.239.91.201:3128 | ✓ 800ms | 否 | ✓ 282ms | ✓ 1469ms | ✓ 881ms | http |
| 35.212.157.53:3128 | ✓ 340ms | 否 | ✓ 990ms | ✓ 1263ms | ✓ 964ms | http |
| 68.62.138.143:3128 | ✓ 930ms | 否 | ✓ 453ms | ✓ 1310ms | ✓ 1054ms | http |
| 84.247.187.120:3128 | ✓ 610ms | ✓ 1828ms | 否 | 否 | ✓ 1441ms | http |
| 5.104.87.17:8050 | ✓ 1459ms | 否 | 否 | ✓ 1284ms | ✓ 1396ms | http |
| 218.108.131.186:17890 | ✓ 877ms | ✓ 1126ms | ✓ 912ms | ✓ 1197ms | ✓ 981ms | http |
| 185.76.240.167:10001 | ✓ 1057ms | 否 | ✓ 1258ms | 否 | ✓ 1805ms | http |
| 185.76.240.169:10001 | ✓ 1115ms | 否 | ✓ 1733ms | 否 | ✓ 1844ms | http |
| 38.145.203.109:8449 | ✓ 1839ms | 否 | ✓ 1779ms | ✓ 977ms | ✓ 1511ms | http |
| 38.34.179.42:8449 | ✓ 1561ms | 否 | ✓ 1867ms | ✓ 1151ms | ✓ 1509ms | http |
| 34.101.184.164:3128 | ✓ 1813ms | 否 | ✓ 1408ms | ✓ 1822ms | 否 | http |
| 104.168.93.120:8080 | ✓ 1257ms | 否 | 否 | ✓ 1329ms | ✓ 1142ms | http |
| 8.219.64.245:3128 | ✓ 1291ms | 否 | ✓ 1012ms | ✓ 1137ms | ✓ 925ms | http |
| 168.222.254.88:3128 | ✓ 727ms | ✓ 1828ms | ✓ 1770ms | ✓ 1664ms | ✓ 1602ms | http |
| 45.128.220.217:3128 | ✓ 1140ms | 否 | 否 | ✓ 887ms | ✓ 770ms | http |
| 171.227.167.109:1008 | ✓ 974ms | 否 | 否 | ✓ 1475ms | ✓ 1052ms | http |
| 150.249.255.91:3128 | ✓ 1503ms | ✓ 1056ms | ✓ 636ms | ✓ 1376ms | 否 | http |
| 150.241.116.228:3128 | ✓ 723ms | 否 | ✓ 971ms | ✓ 1936ms | 否 | http |
| 45.136.130.188:8449 | ✓ 940ms | ✓ 1700ms | 否 | ✓ 1873ms | 否 | http |
| 110.42.37.202:20005 | ✓ 1782ms | ✓ 1767ms | ✓ 1726ms | 否 | ✓ 1347ms | http |
| 38.145.208.209:8447 | ✓ 1737ms | ✓ 941ms | 否 | ✓ 1747ms | ✓ 1736ms | http |
| 38.34.179.177:8446 | ✓ 1356ms | 否 | ✓ 672ms | ✓ 1031ms | 否 | http |
| 38.34.179.176:8446 | ✓ 1357ms | 否 | ✓ 670ms | ✓ 1014ms | 否 | http |
| 38.145.208.224:8445 | ✓ 1421ms | ✓ 1857ms | ✓ 1233ms | ✓ 1975ms | ✓ 900ms | http |
| 38.145.208.245:8449 | 否 | ✓ 1049ms | ✓ 814ms | ✓ 1852ms | ✓ 1011ms | http |
| 38.34.179.38:8447 | ✓ 1531ms | 否 | ✓ 646ms | 否 | ✓ 1884ms | http |
| 121.230.9.101:1080 | 否 | ✓ 1544ms | ✓ 1159ms | ✓ 1386ms | ✓ 1207ms | http |
| 181.78.44.63:999 | ✓ 1452ms | 否 | ✓ 504ms | ✓ 1435ms | 否 | http |
| 185.76.240.254:10001 | ✓ 1093ms | 否 | ✓ 1466ms | 否 | ✓ 1958ms | http |
| 139.159.99.242:8080 | ✓ 1107ms | ✓ 1098ms | ✓ 1084ms | ✓ 1141ms | 否 | http |
| 38.145.208.220:8448 | ✓ 1984ms | ✓ 965ms | ✓ 949ms | 否 | ✓ 687ms | http |
| 38.145.208.210:8448 | ✓ 459ms | 否 | ✓ 688ms | ✓ 835ms | ✓ 980ms | http |
| 38.145.208.182:8452 | ✓ 1680ms | ✓ 959ms | ✓ 899ms | ✓ 1729ms | ✓ 924ms | http |
| 38.145.218.216:8449 | ✓ 1770ms | ✓ 1577ms | 否 | ✓ 910ms | ✓ 1294ms | http |
| 38.145.218.232:8448 | ✓ 468ms | ✓ 980ms | ✓ 1849ms | ✓ 1895ms | ✓ 966ms | http |
| 38.145.203.39:8445 | ✓ 1161ms | 否 | ✓ 1742ms | 否 | ✓ 852ms | http |
| 147.161.246.38:11814 | ✓ 1911ms | 否 | ✓ 1657ms | 否 | ✓ 1564ms | http |

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
