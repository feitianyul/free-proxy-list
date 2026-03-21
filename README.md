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

最后更新：2026-03-21 17:27:46 UTC（2026-03-22 01:27:46 UTC+8）

**代理总数：81**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | 否 | 否 | ✓ 1170ms | ✓ 1146ms | ✓ 875ms | http |
| 167.103.34.108:8800 | ✓ 1327ms | 否 | 否 | ✓ 1494ms | ✓ 1389ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1637ms | 否 | ✓ 1965ms | ✓ 1623ms | http |
| 113.160.132.26:8080 | ✓ 1600ms | 否 | 否 | ✓ 1409ms | ✓ 1106ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1854ms | ✓ 1512ms | ✓ 1685ms | ✓ 1471ms | http |
| 103.113.70.189:1081 | ✓ 293ms | 否 | ✓ 1019ms | ✓ 1537ms | ✓ 1024ms | http |
| 133.242.138.34:8100 | ✓ 696ms | 否 | ✓ 695ms | 否 | ✓ 857ms | http |
| 103.84.95.54:7890 | ✓ 885ms | 否 | ✓ 1006ms | ✓ 1864ms | 否 | http |
| 137.220.151.110:6005 | ✓ 1074ms | 否 | ✓ 896ms | 否 | ✓ 1297ms | http |
| 2.56.173.45:10808 | ✓ 1141ms | 否 | ✓ 826ms | 否 | ✓ 1637ms | http |
| 167.103.31.122:8800 | ✓ 1625ms | 否 | ✓ 1834ms | 否 | ✓ 1528ms | http |
| 167.71.60.190:8080 | ✓ 509ms | 否 | ✓ 1642ms | ✓ 1836ms | 否 | http |
| 137.220.150.22:6005 | ✓ 1562ms | 否 | ✓ 966ms | ✓ 1347ms | ✓ 1042ms | http |
| 91.233.223.147:3128 | ✓ 1473ms | 否 | 否 | ✓ 1888ms | ✓ 1959ms | http |
| 223.16.170.103:80 | ✓ 1028ms | ✓ 1884ms | ✓ 1184ms | ✓ 1219ms | 否 | http |
| 35.225.22.61:80 | ✓ 1046ms | ✓ 1705ms | 否 | 否 | ✓ 983ms | http |
| 38.145.208.172:8448 | ✓ 1082ms | ✓ 919ms | ✓ 408ms | ✓ 1427ms | ✓ 1529ms | http |
| 104.168.158.236:10808 | ✓ 1041ms | ✓ 1549ms | ✓ 1797ms | ✓ 1355ms | ✓ 1747ms | http |
| 147.161.239.240:8800 | ✓ 1224ms | ✓ 1712ms | ✓ 1300ms | ✓ 1601ms | ✓ 1567ms | http |
| 192.71.213.85:9812 | ✓ 1224ms | 否 | ✓ 1595ms | ✓ 1711ms | 否 | http |
| 45.136.130.174:8450 | ✓ 1326ms | ✓ 1415ms | ✓ 1044ms | 否 | ✓ 1339ms | http |
| 91.238.105.64:2024 | ✓ 1249ms | ✓ 1919ms | ✓ 1874ms | 否 | ✓ 1736ms | http |
| 120.92.212.16:8890 | ✓ 1056ms | ✓ 1351ms | ✓ 1105ms | ✓ 1398ms | 否 | http |
| 38.145.208.226:8453 | ✓ 749ms | 否 | ✓ 320ms | ✓ 1758ms | 否 | http |
| 38.145.208.172:8453 | ✓ 944ms | 否 | ✓ 485ms | ✓ 1578ms | ✓ 1785ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1471ms | ✓ 1486ms | ✓ 1701ms | ✓ 1490ms | http |
| 219.117.204.211:7799 | ✓ 1627ms | 否 | ✓ 1144ms | 否 | ✓ 1075ms | http |
| 140.82.34.41:1080 | ✓ 1061ms | 否 | ✓ 1349ms | 否 | ✓ 1878ms | http |
| 77.110.113.24:40000 | ✓ 1056ms | 否 | ✓ 1397ms | ✓ 1992ms | 否 | http |
| 101.43.127.100:8877 | 否 | ✓ 1290ms | 否 | ✓ 1275ms | ✓ 1299ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1817ms | ✓ 1339ms | ✓ 1654ms | 否 | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1588ms | ✓ 1795ms | ✓ 1955ms | http |
| 86.109.3.24:10008 | 否 | ✓ 1860ms | ✓ 446ms | ✓ 793ms | ✓ 616ms | http |
| 86.109.3.24:10005 | ✓ 180ms | ✓ 1884ms | ✓ 270ms | ✓ 921ms | ✓ 612ms | http |
| 86.109.3.24:10012 | ✓ 179ms | 否 | ✓ 272ms | ✓ 804ms | ✓ 623ms | http |
| 139.162.46.62:3128 | ✓ 1455ms | 否 | 否 | ✓ 1177ms | ✓ 1577ms | http |
| 83.219.250.8:62920 | ✓ 946ms | 否 | ✓ 1473ms | 否 | ✓ 1468ms | http |
| 194.67.99.223:1080 | ✓ 966ms | 否 | 否 | ✓ 1910ms | ✓ 1383ms | http |
| 142.171.224.229:7890 | ✓ 565ms | ✓ 827ms | ✓ 341ms | ✓ 1191ms | ✓ 599ms | http |
| 137.220.150.170:6005 | ✓ 1947ms | 否 | ✓ 1803ms | 否 | ✓ 1196ms | http |
| 150.241.77.172:1080 | ✓ 860ms | ✓ 1903ms | ✓ 442ms | ✓ 1782ms | 否 | http |
| 43.165.195.107:3128 | 否 | ✓ 1956ms | ✓ 1054ms | ✓ 1330ms | ✓ 1048ms | http |
| 45.136.130.171:8445 | ✓ 1171ms | ✓ 1116ms | ✓ 990ms | ✓ 1523ms | ✓ 901ms | http |
| 193.23.200.251:10808 | ✓ 1141ms | ✓ 1789ms | ✓ 1545ms | 否 | ✓ 1334ms | http |
| 38.34.179.61:8445 | ✓ 579ms | ✓ 1919ms | ✓ 1995ms | ✓ 881ms | ✓ 1542ms | http |
| 38.34.179.96:8451 | ✓ 257ms | ✓ 1762ms | 否 | ✓ 945ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1720ms | 否 | 否 | ✓ 1963ms | ✓ 1600ms | https |
| 38.145.208.239:8446 | ✓ 376ms | ✓ 1144ms | ✓ 968ms | ✓ 1010ms | ✓ 1007ms | http |
| 38.34.179.35:8453 | ✓ 287ms | 否 | ✓ 1029ms | ✓ 952ms | ✓ 1103ms | http |
| 38.145.218.161:8445 | ✓ 1295ms | ✓ 1239ms | ✓ 888ms | ✓ 1878ms | ✓ 732ms | http |
| 38.145.218.101:8448 | ✓ 1824ms | 否 | ✓ 1292ms | ✓ 864ms | ✓ 860ms | http |
| 144.31.79.117:8888 | ✓ 682ms | 否 | ✓ 1321ms | 否 | ✓ 1724ms | http |
| 38.145.208.244:8452 | ✓ 1899ms | ✓ 1725ms | 否 | ✓ 977ms | ✓ 1462ms | http |
| 38.145.208.227:8451 | ✓ 762ms | 否 | ✓ 529ms | 否 | ✓ 1530ms | http |
| 137.220.150.104:6005 | ✓ 1182ms | 否 | 否 | ✓ 1504ms | ✓ 1152ms | http |
| 150.107.140.238:3128 | ✓ 1817ms | 否 | ✓ 1373ms | ✓ 1427ms | 否 | http |
| 106.75.15.167:7890 | ✓ 1260ms | 否 | ✓ 1579ms | ✓ 1270ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1298ms | ✓ 1709ms | ✓ 1846ms | ✓ 1042ms | ✓ 813ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1908ms | ✓ 1887ms | ✓ 1125ms | http |
| 166.88.55.83:7890 | ✓ 741ms | ✓ 1634ms | ✓ 748ms | ✓ 950ms | ✓ 750ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1355ms | ✓ 1105ms | ✓ 1692ms | 否 | http |
| 45.136.130.186:8451 | ✓ 692ms | ✓ 1729ms | ✓ 1338ms | ✓ 893ms | ✓ 715ms | http |
| 103.183.10.169:3125 | ✓ 1514ms | 否 | ✓ 1775ms | ✓ 1818ms | ✓ 1641ms | http |
| 212.192.12.90:6005 | 否 | 否 | ✓ 1881ms | ✓ 1847ms | ✓ 1528ms | http |
| 45.186.6.104:3128 | ✓ 1193ms | ✓ 1958ms | ✓ 1526ms | 否 | 否 | http |
| 172.212.68.37:3128 | ✓ 272ms | ✓ 1588ms | 否 | ✓ 1320ms | ✓ 1148ms | http |
| 209.97.149.157:80 | ✓ 727ms | ✓ 1930ms | ✓ 1652ms | ✓ 1324ms | ✓ 1384ms | http |
| 45.136.198.40:3128 | ✓ 1283ms | ✓ 1962ms | ✓ 1822ms | 否 | ✓ 1785ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1735ms | ✓ 1318ms | ✓ 1237ms | 否 | http |
| 103.82.23.118:5242 | ✓ 1594ms | 否 | 否 | ✓ 1657ms | ✓ 1611ms | http |
| 45.140.147.155:1081 | ✓ 912ms | ✓ 1136ms | ✓ 1142ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 543ms | ✓ 1977ms | ✓ 463ms | 否 | 否 | http |
| 125.64.244.100:8889 | ✓ 1808ms | ✓ 1873ms | ✓ 1790ms | 否 | 否 | http |
| 34.96.238.40:8080 | ✓ 944ms | 否 | ✓ 1487ms | 否 | ✓ 1516ms | http |
| 121.126.185.63:25152 | ✓ 1971ms | 否 | 否 | ✓ 1916ms | ✓ 1574ms | http |
| 61.52.131.172:8443 | ✓ 1028ms | ✓ 1299ms | ✓ 996ms | ✓ 1320ms | ✓ 984ms | http |
| 38.34.178.186:8451 | ✓ 801ms | ✓ 1329ms | ✓ 1282ms | ✓ 1645ms | ✓ 666ms | http |
| 101.47.73.135:3128 | ✓ 1842ms | 否 | ✓ 1892ms | 否 | ✓ 1372ms | http |
| 202.141.161.53:30001 | ✓ 1214ms | ✓ 1479ms | ✓ 1598ms | 否 | 否 | http |
| 4.233.138.204:8888 | ✓ 1534ms | ✓ 1957ms | 否 | 否 | ✓ 1812ms | http |
| 103.39.51.190:8080 | ✓ 1932ms | 否 | 否 | ✓ 1680ms | ✓ 1588ms | http |

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
