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

最后更新：2026-03-21 16:28:14 UTC（2026-03-22 00:28:14 UTC+8）

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
| 147.161.210.140:8800 | ✓ 1442ms | 否 | ✓ 971ms | ✓ 1125ms | ✓ 1070ms | http |
| 139.159.99.242:8080 | ✓ 918ms | ✓ 1166ms | ✓ 928ms | ✓ 1215ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1640ms | 否 | ✓ 1365ms | 否 | ✓ 1108ms | http |
| 38.145.208.185:8449 | ✓ 277ms | ✓ 1693ms | ✓ 707ms | ✓ 1254ms | ✓ 996ms | http |
| 219.117.204.211:7799 | ✓ 690ms | 否 | ✓ 651ms | ✓ 970ms | ✓ 820ms | http |
| 38.145.208.181:8445 | ✓ 906ms | ✓ 814ms | ✓ 930ms | ✓ 1319ms | ✓ 1157ms | http |
| 104.168.158.236:10808 | ✓ 916ms | 否 | 否 | ✓ 924ms | ✓ 812ms | http |
| 38.34.179.40:8446 | ✓ 310ms | ✓ 840ms | ✓ 1148ms | 否 | ✓ 982ms | http |
| 38.34.179.172:8451 | ✓ 598ms | ✓ 1427ms | ✓ 378ms | 否 | ✓ 958ms | http |
| 38.34.179.165:8446 | ✓ 612ms | ✓ 1517ms | ✓ 1111ms | 否 | ✓ 1048ms | http |
| 133.242.138.34:8100 | ✓ 735ms | ✓ 1453ms | ✓ 1166ms | ✓ 1111ms | 否 | http |
| 137.220.150.152:6005 | ✓ 1072ms | 否 | ✓ 1348ms | ✓ 1572ms | ✓ 1196ms | http |
| 167.103.34.108:8800 | ✓ 1286ms | 否 | ✓ 1421ms | ✓ 1420ms | ✓ 1309ms | http |
| 38.145.208.179:8447 | ✓ 1223ms | 否 | ✓ 1580ms | ✓ 1572ms | 否 | http |
| 38.145.203.19:8449 | ✓ 1968ms | ✓ 1256ms | ✓ 302ms | ✓ 1479ms | 否 | http |
| 45.136.131.54:8448 | ✓ 776ms | 否 | ✓ 1618ms | 否 | ✓ 1797ms | http |
| 120.92.212.16:7890 | ✓ 1055ms | 否 | 否 | ✓ 1633ms | ✓ 1064ms | http |
| 120.92.212.16:8890 | ✓ 1069ms | 否 | 否 | ✓ 1319ms | ✓ 1327ms | http |
| 167.103.31.122:8800 | ✓ 1579ms | 否 | ✓ 1562ms | ✓ 1622ms | ✓ 1603ms | http |
| 210.76.193.248:10808 | ✓ 1348ms | 否 | 否 | ✓ 1459ms | ✓ 1484ms | http |
| 172.212.68.37:3128 | ✓ 229ms | 否 | 否 | ✓ 1318ms | ✓ 1074ms | http |
| 147.161.239.240:8800 | ✓ 1084ms | 否 | ✓ 1321ms | ✓ 1650ms | ✓ 1342ms | http |
| 91.238.105.64:2024 | ✓ 1114ms | 否 | 否 | ✓ 1934ms | ✓ 1833ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1982ms | ✓ 1441ms | 否 | ✓ 1296ms | http |
| 137.220.150.22:6005 | ✓ 872ms | 否 | ✓ 1050ms | 否 | ✓ 1024ms | http |
| 91.233.223.147:3128 | ✓ 871ms | 否 | ✓ 1169ms | ✓ 1954ms | ✓ 1513ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1174ms | ✓ 1692ms | ✓ 1476ms | ✓ 1256ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1871ms | ✓ 726ms | ✓ 1023ms | 否 | http |
| 150.249.255.91:3128 | ✓ 1388ms | 否 | ✓ 677ms | 否 | ✓ 1009ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1407ms | ✓ 1989ms | ✓ 1689ms | 否 | http |
| 171.227.173.219:10001 | 否 | 否 | ✓ 1711ms | ✓ 1852ms | ✓ 1611ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 255ms | ✓ 959ms | ✓ 778ms | http |
| 137.220.151.110:6005 | ✓ 923ms | 否 | ✓ 1093ms | ✓ 1413ms | 否 | http |
| 38.34.179.173:8452 | ✓ 896ms | ✓ 935ms | ✓ 417ms | ✓ 1353ms | 否 | http |
| 16.78.119.130:443 | ✓ 1799ms | 否 | ✓ 1581ms | ✓ 1889ms | 否 | http |
| 38.34.179.39:8452 | ✓ 914ms | ✓ 961ms | ✓ 909ms | ✓ 1248ms | 否 | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1097ms | ✓ 1327ms | ✓ 1729ms | http |
| 38.34.179.74:8449 | ✓ 1407ms | 否 | ✓ 295ms | ✓ 1290ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1855ms | 否 | 否 | ✓ 1354ms | ✓ 1077ms | http |
| 59.46.216.131:30001 | ✓ 1432ms | 否 | ✓ 1919ms | ✓ 1982ms | 否 | http |
| 38.145.218.229:8450 | ✓ 472ms | ✓ 1714ms | ✓ 409ms | ✓ 1246ms | 否 | http |
| 144.31.79.117:8888 | ✓ 1031ms | 否 | ✓ 1445ms | ✓ 1743ms | ✓ 1393ms | http |
| 45.136.130.178:8449 | 否 | ✓ 1642ms | ✓ 221ms | ✓ 1260ms | ✓ 1136ms | http |
| 194.67.99.223:1080 | ✓ 1011ms | 否 | ✓ 1708ms | 否 | ✓ 1430ms | http |
| 38.145.208.182:8453 | ✓ 1770ms | ✓ 820ms | ✓ 419ms | 否 | ✓ 1438ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1150ms | ✓ 1459ms | ✓ 1546ms | http |
| 38.145.208.237:8443 | ✓ 1992ms | 否 | ✓ 701ms | 否 | ✓ 958ms | http |
| 38.145.208.235:8443 | 否 | ✓ 1925ms | ✓ 362ms | 否 | ✓ 1943ms | http |
| 192.71.213.85:9812 | ✓ 1099ms | 否 | ✓ 1899ms | ✓ 1775ms | 否 | http |
| 24.199.124.152:3128 | ✓ 433ms | ✓ 1554ms | ✓ 1121ms | ✓ 868ms | ✓ 1882ms | http |
| 164.90.155.209:3128 | ✓ 1168ms | ✓ 988ms | ✓ 393ms | ✓ 1026ms | ✓ 650ms | http |
| 142.171.224.229:7890 | ✓ 1186ms | ✓ 846ms | ✓ 908ms | ✓ 1235ms | ✓ 610ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1942ms | ✓ 1762ms | ✓ 1612ms | http |
| 45.136.130.185:8444 | ✓ 1924ms | ✓ 1673ms | 否 | ✓ 1595ms | 否 | http |
| 38.34.179.35:8453 | ✓ 334ms | ✓ 1495ms | ✓ 397ms | ✓ 1414ms | ✓ 1113ms | http |
| 20.27.14.220:8561 | ✓ 585ms | ✓ 996ms | ✓ 732ms | ✓ 949ms | ✓ 737ms | http |
| 38.34.179.34:8449 | ✓ 423ms | ✓ 1013ms | ✓ 789ms | ✓ 1432ms | ✓ 1101ms | http |
| 38.145.208.227:8451 | ✓ 561ms | 否 | ✓ 867ms | ✓ 1451ms | ✓ 1101ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1680ms | 否 | ✓ 1348ms | ✓ 1444ms | http |
| 45.136.130.179:8444 | ✓ 319ms | 否 | ✓ 280ms | ✓ 1317ms | 否 | http |
| 38.34.179.213:8452 | ✓ 410ms | 否 | ✓ 562ms | ✓ 1279ms | 否 | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 602ms | ✓ 920ms | ✓ 717ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 586ms | ✓ 905ms | ✓ 725ms | http |
| 20.27.15.111:8561 | 否 | ✓ 859ms | ✓ 586ms | ✓ 947ms | ✓ 739ms | http |
| 166.88.55.83:7890 | ✓ 742ms | ✓ 1629ms | ✓ 735ms | ✓ 943ms | ✓ 766ms | http |
| 202.38.72.235:26001 | ✓ 1792ms | 否 | ✓ 1960ms | 否 | ✓ 1850ms | http |
| 38.34.179.54:8446 | ✓ 556ms | 否 | ✓ 292ms | ✓ 1687ms | 否 | http |
| 45.149.92.147:5001 | ✓ 1489ms | 否 | ✓ 978ms | ✓ 1374ms | ✓ 780ms | http |
| 38.145.208.242:8451 | ✓ 871ms | 否 | ✓ 1358ms | ✓ 1714ms | 否 | http |
| 23.148.244.54:20046 | 否 | 否 | ✓ 373ms | ✓ 1186ms | ✓ 1016ms | http |
| 2.56.173.45:10808 | ✓ 1371ms | 否 | ✓ 1427ms | 否 | ✓ 1578ms | http |
| 1.231.81.166:3128 | ✓ 1046ms | ✓ 1069ms | 否 | ✓ 1338ms | ✓ 1040ms | http |
| 38.34.179.98:8453 | 否 | ✓ 920ms | ✓ 675ms | ✓ 1718ms | ✓ 1636ms | http |
| 223.16.170.103:80 | ✓ 1263ms | 否 | ✓ 1408ms | ✓ 1228ms | ✓ 1679ms | http |
| 38.145.208.172:8448 | ✓ 749ms | ✓ 845ms | ✓ 437ms | ✓ 1248ms | ✓ 1042ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 849ms | ✓ 1883ms | ✓ 779ms | http |
| 152.69.229.220:3128 | 否 | 否 | ✓ 1732ms | ✓ 1570ms | ✓ 1100ms | http |
| 45.136.198.40:3128 | ✓ 1137ms | 否 | ✓ 1870ms | 否 | ✓ 1934ms | http |
| 101.230.73.57:29999 | ✓ 926ms | 否 | ✓ 973ms | 否 | ✓ 946ms | http |
| 61.52.131.172:8443 | ✓ 1026ms | ✓ 1280ms | ✓ 991ms | ✓ 1282ms | ✓ 1028ms | http |
| 45.136.130.168:8448 | ✓ 399ms | ✓ 928ms | ✓ 422ms | 否 | ✓ 1174ms | http |
| 45.136.130.170:8448 | ✓ 1174ms | 否 | ✓ 356ms | ✓ 1383ms | ✓ 1676ms | http |
| 49.156.44.114:8080 | ✓ 1729ms | ✓ 1872ms | ✓ 1829ms | ✓ 1700ms | ✓ 1754ms | http |
| 106.117.208.101:7890 | ✓ 1053ms | ✓ 1394ms | ✓ 1172ms | ✓ 1424ms | ✓ 1164ms | http |
| 38.34.179.61:8445 | ✓ 990ms | ✓ 1265ms | ✓ 674ms | ✓ 1952ms | ✓ 1663ms | http |
| 38.34.179.178:8445 | ✓ 617ms | 否 | ✓ 1424ms | 否 | ✓ 1928ms | http |
| 137.184.6.37:3128 | ✓ 1443ms | 否 | 否 | ✓ 977ms | ✓ 647ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 299ms | ✓ 1351ms | ✓ 1033ms | http |

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
