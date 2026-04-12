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

最后更新：2026-04-12 10:42:50 UTC（2026-04-12 18:42:50 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 854ms | 否 | ✓ 1645ms | ✓ 1214ms | ✓ 1113ms | http |
| 185.198.27.38:3128 | ✓ 888ms | ✓ 1671ms | ✓ 1649ms | ✓ 1725ms | ✓ 1356ms | http |
| 167.103.115.102:8800 | ✓ 1957ms | 否 | ✓ 1336ms | ✓ 1461ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1647ms | ✓ 1607ms | ✓ 1852ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1267ms | 否 | ✓ 1467ms | ✓ 1783ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1623ms | 否 | ✓ 1451ms | ✓ 1603ms | ✓ 1624ms | http |
| 137.59.47.73:3128 | ✓ 1261ms | 否 | ✓ 1728ms | 否 | ✓ 1500ms | http |
| 45.167.125.21:999 | ✓ 1290ms | ✓ 1844ms | ✓ 1221ms | ✓ 1991ms | ✓ 1670ms | http |
| 167.103.31.122:8800 | ✓ 1527ms | 否 | ✓ 1293ms | ✓ 1599ms | ✓ 1443ms | http |
| 116.80.63.178:3172 | ✓ 1719ms | 否 | ✓ 1623ms | ✓ 1933ms | 否 | http |
| 147.161.239.240:8800 | ✓ 571ms | ✓ 1659ms | ✓ 1044ms | ✓ 1559ms | ✓ 1377ms | http |
| 46.30.46.133:3128 | ✓ 1078ms | 否 | 否 | ✓ 1308ms | ✓ 1301ms | http |
| 159.223.225.118:8888 | ✓ 1661ms | 否 | ✓ 1469ms | 否 | ✓ 1130ms | http |
| 38.34.179.178:8444 | ✓ 870ms | 否 | ✓ 1714ms | ✓ 1041ms | ✓ 1318ms | http |
| 202.141.161.53:10808 | ✓ 1138ms | ✓ 1677ms | ✓ 1867ms | 否 | ✓ 1267ms | http |
| 35.225.22.61:80 | ✓ 485ms | ✓ 1306ms | ✓ 256ms | ✓ 1044ms | 否 | http |
| 121.230.9.203:1080 | ✓ 1320ms | ✓ 1614ms | ✓ 1568ms | ✓ 1687ms | ✓ 1262ms | http |
| 54.222.174.194:80 | 否 | ✓ 1774ms | ✓ 1598ms | ✓ 1911ms | 否 | http |
| 85.239.59.252:7890 | ✓ 1236ms | ✓ 1563ms | 否 | ✓ 1789ms | ✓ 1315ms | http |
| 46.39.105.157:8080 | ✓ 573ms | ✓ 1634ms | ✓ 1817ms | ✓ 1602ms | ✓ 1194ms | http |
| 34.231.145.203:7000 | ✓ 107ms | ✓ 925ms | ✓ 782ms | ✓ 952ms | ✓ 1912ms | http |
| 177.93.132.244:3128 | ✓ 1094ms | 否 | ✓ 889ms | 否 | ✓ 1615ms | http |
| 121.230.9.125:1080 | ✓ 1130ms | 否 | ✓ 1345ms | ✓ 1566ms | ✓ 1774ms | http |
| 223.84.151.86:30005 | ✓ 1390ms | ✓ 1453ms | ✓ 1246ms | ✓ 1572ms | ✓ 1207ms | http |
| 110.42.37.202:20005 | ✓ 1226ms | 否 | ✓ 1263ms | ✓ 1683ms | ✓ 1365ms | http |
| 120.92.108.86:7890 | ✓ 1724ms | 否 | ✓ 1816ms | 否 | ✓ 1945ms | http |
| 94.72.109.214:8888 | ✓ 1164ms | 否 | ✓ 990ms | 否 | ✓ 1667ms | http |
| 103.231.75.209:3128 | ✓ 1653ms | 否 | ✓ 1051ms | 否 | ✓ 1574ms | http |
| 162.240.154.26:3128 | ✓ 1092ms | 否 | ✓ 1270ms | ✓ 1718ms | ✓ 1129ms | http |
| 103.125.56.83:8080 | ✓ 1658ms | 否 | ✓ 1506ms | ✓ 1707ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1321ms | 否 | ✓ 1583ms | ✓ 1578ms | ✓ 1291ms | http |
| 103.165.157.206:8088 | ✓ 1473ms | 否 | ✓ 1379ms | ✓ 1594ms | ✓ 1565ms | http |
| 103.134.245.127:8090 | ✓ 1628ms | 否 | 否 | ✓ 1770ms | ✓ 1957ms | http |
| 38.242.212.16:3128 | ✓ 1575ms | ✓ 1372ms | ✓ 813ms | ✓ 1543ms | ✓ 1221ms | http |
| 185.76.240.117:10001 | ✓ 752ms | 否 | ✓ 1274ms | 否 | ✓ 1842ms | http |
| 5.196.101.18:3128 | 否 | 否 | ✓ 735ms | ✓ 1753ms | ✓ 1448ms | http |
| 197.164.101.11:1976 | ✓ 1594ms | 否 | ✓ 1546ms | 否 | ✓ 1789ms | http |
| 115.231.181.40:8128 | ✓ 1077ms | ✓ 1585ms | ✓ 1118ms | ✓ 1443ms | ✓ 1098ms | http |
| 36.141.21.200:7890 | ✓ 1134ms | ✓ 1477ms | ✓ 1174ms | ✓ 1438ms | ✓ 1076ms | http |
| 36.103.198.235:7890 | ✓ 1187ms | ✓ 1657ms | ✓ 1252ms | ✓ 1869ms | ✓ 1151ms | http |
| 5.255.123.43:1080 | ✓ 530ms | 否 | 否 | ✓ 1734ms | ✓ 1082ms | http |
| 185.76.241.110:10001 | ✓ 980ms | 否 | ✓ 1744ms | 否 | ✓ 1946ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1313ms | ✓ 1448ms | ✓ 1606ms | http |
| 139.159.99.242:8080 | ✓ 920ms | 否 | ✓ 961ms | ✓ 1217ms | ✓ 1015ms | http |
| 62.234.206.73:3128 | 否 | ✓ 1536ms | ✓ 1069ms | ✓ 1770ms | ✓ 1805ms | http |
| 180.130.80.196:9003 | ✓ 1474ms | ✓ 1506ms | 否 | 否 | ✓ 1280ms | http |
| 185.76.240.167:10001 | ✓ 1468ms | 否 | ✓ 1357ms | 否 | ✓ 1742ms | http |
| 109.107.179.140:31000 | ✓ 1472ms | 否 | ✓ 1281ms | 否 | ✓ 1270ms | http |
| 130.162.53.123:50960 | ✓ 790ms | ✓ 1423ms | ✓ 1499ms | ✓ 1317ms | 否 | http |
| 133.18.110.87:1081 | ✓ 1696ms | 否 | 否 | ✓ 1126ms | ✓ 874ms | http |
| 38.34.179.105:8447 | 否 | ✓ 1974ms | ✓ 1720ms | ✓ 1461ms | ✓ 779ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1383ms | ✓ 1874ms | 否 | ✓ 1148ms | http |
| 121.230.8.220:1080 | ✓ 1443ms | ✓ 1460ms | ✓ 1233ms | 否 | 否 | http |
| 121.230.8.45:1080 | ✓ 1173ms | ✓ 1577ms | ✓ 1608ms | ✓ 1718ms | ✓ 1122ms | http |
| 121.230.8.39:1080 | ✓ 1182ms | ✓ 1579ms | ✓ 1439ms | ✓ 1833ms | ✓ 1488ms | http |
| 181.78.44.63:999 | ✓ 585ms | ✓ 1545ms | ✓ 1193ms | ✓ 1693ms | ✓ 1210ms | http |
| 45.140.147.155:1082 | ✓ 661ms | 否 | ✓ 691ms | ✓ 1762ms | 否 | http |
| 38.145.218.161:8445 | ✓ 927ms | ✓ 999ms | ✓ 1680ms | ✓ 1702ms | 否 | http |
| 45.136.130.169:8444 | ✓ 935ms | 否 | ✓ 1379ms | ✓ 1057ms | 否 | http |
| 38.34.179.8:8443 | ✓ 949ms | 否 | ✓ 523ms | ✓ 1126ms | ✓ 774ms | http |
| 38.145.208.211:8453 | ✓ 1310ms | ✓ 1201ms | ✓ 454ms | ✓ 923ms | ✓ 969ms | http |
| 38.34.179.40:8446 | ✓ 1259ms | 否 | ✓ 728ms | ✓ 1404ms | ✓ 763ms | http |
| 38.145.208.169:8452 | ✓ 925ms | ✓ 1090ms | ✓ 1109ms | 否 | ✓ 815ms | http |
| 38.145.208.209:8447 | ✓ 1865ms | ✓ 1070ms | ✓ 537ms | 否 | ✓ 982ms | http |
| 45.136.130.166:8447 | ✓ 717ms | 否 | ✓ 745ms | ✓ 1039ms | ✓ 820ms | http |
| 38.34.179.79:8451 | ✓ 670ms | 否 | ✓ 1066ms | ✓ 1006ms | ✓ 909ms | http |
| 38.145.220.8:8452 | ✓ 1727ms | ✓ 1192ms | 否 | ✓ 1659ms | ✓ 1118ms | http |
| 38.34.183.221:8452 | ✓ 343ms | 否 | ✓ 1728ms | ✓ 1004ms | ✓ 1828ms | http |
| 5.104.87.17:8051 | ✓ 1919ms | 否 | ✓ 1944ms | ✓ 1429ms | ✓ 909ms | http |
| 223.16.170.103:80 | ✓ 1654ms | 否 | ✓ 1398ms | ✓ 1274ms | ✓ 1269ms | http |
| 38.145.218.134:8446 | ✓ 703ms | ✓ 1740ms | ✓ 413ms | ✓ 1069ms | ✓ 1059ms | http |
| 38.34.179.33:8450 | ✓ 785ms | ✓ 1076ms | 否 | ✓ 1185ms | ✓ 1670ms | http |
| 38.145.220.103:8446 | 否 | ✓ 1818ms | ✓ 1115ms | ✓ 1658ms | ✓ 994ms | http |

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
