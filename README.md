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

最后更新：2026-02-28 12:19:08 UTC（2026-02-28 20:19:08 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 130ms | 否 | ✓ 668ms | ✓ 1094ms | ✓ 812ms | http |
| 72.56.59.62:63133 | ✓ 1481ms | 否 | ✓ 1742ms | 否 | ✓ 1874ms | http |
| 104.238.30.40:59741 | ✓ 1603ms | 否 | ✓ 1998ms | 否 | ✓ 1867ms | http |
| 104.238.30.37:59741 | ✓ 1581ms | 否 | ✓ 1679ms | 否 | ✓ 1907ms | http |
| 120.92.212.16:8890 | ✓ 1001ms | ✓ 1308ms | ✓ 1125ms | ✓ 1321ms | ✓ 1047ms | http |
| 142.171.85.32:1080 | ✓ 611ms | 否 | ✓ 1071ms | ✓ 1034ms | ✓ 1040ms | http |
| 81.177.48.54:2080 | ✓ 1613ms | 否 | ✓ 1653ms | ✓ 1743ms | 否 | http |
| 104.238.30.45:59741 | ✓ 1608ms | 否 | ✓ 1936ms | 否 | ✓ 1871ms | http |
| 115.231.181.40:8128 | ✓ 1956ms | ✓ 1962ms | 否 | ✓ 1244ms | ✓ 1544ms | http |
| 104.238.30.58:63744 | ✓ 1568ms | 否 | ✓ 1967ms | 否 | ✓ 1871ms | http |
| 35.234.17.221:8080 | ✓ 1026ms | ✓ 1415ms | ✓ 1146ms | 否 | 否 | http |
| 3.213.157.4:3128 | ✓ 698ms | ✓ 1534ms | ✓ 1109ms | 否 | ✓ 1899ms | http |
| 104.238.30.50:59741 | ✓ 1609ms | 否 | ✓ 1711ms | 否 | ✓ 1871ms | http |
| 104.238.30.63:63744 | ✓ 1548ms | 否 | ✓ 1771ms | 否 | ✓ 1843ms | http |
| 35.225.22.61:80 | ✓ 862ms | ✓ 1442ms | ✓ 1075ms | ✓ 1208ms | ✓ 695ms | http |
| 85.208.108.43:2094 | ✓ 1053ms | 否 | ✓ 1009ms | ✓ 1125ms | ✓ 909ms | http |
| 120.92.212.16:7890 | ✓ 1016ms | ✓ 1554ms | ✓ 1044ms | ✓ 1589ms | ✓ 1295ms | http |
| 72.56.59.56:63127 | ✓ 1927ms | 否 | ✓ 1906ms | 否 | ✓ 1835ms | http |
| 59.46.216.131:30001 | ✓ 1301ms | ✓ 1461ms | 否 | ✓ 1659ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1288ms | 否 | 否 | ✓ 1385ms | ✓ 1287ms | http |
| 168.235.110.63:3128 | ✓ 205ms | 否 | ✓ 1223ms | 否 | ✓ 1993ms | http |
| 165.227.5.10:8888 | ✓ 1818ms | ✓ 1404ms | 否 | ✓ 1120ms | ✓ 906ms | http |
| 36.147.78.166:80 | ✓ 1898ms | 否 | 否 | ✓ 1906ms | ✓ 1880ms | http |
| 81.70.169.194:80 | 否 | ✓ 1386ms | ✓ 1206ms | ✓ 1368ms | 否 | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1678ms | ✓ 1643ms | ✓ 1549ms | http |
| 101.43.255.96:80 | ✓ 1930ms | ✓ 1431ms | ✓ 1494ms | ✓ 1946ms | ✓ 1414ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1392ms | ✓ 943ms | ✓ 1303ms | ✓ 1188ms | http |
| 85.208.108.43:10808 | ✓ 409ms | 否 | ✓ 331ms | ✓ 1089ms | ✓ 893ms | http |
| 62.113.119.14:8080 | ✓ 1836ms | 否 | ✓ 987ms | ✓ 1607ms | 否 | http |
| 103.133.254.4:3128 | ✓ 1375ms | 否 | ✓ 1902ms | ✓ 1792ms | ✓ 1468ms | http |
| 147.45.216.148:1080 | ✓ 870ms | 否 | 否 | ✓ 1861ms | ✓ 1256ms | http |
| 103.215.36.88:17853 | 否 | ✓ 1420ms | ✓ 1273ms | 否 | ✓ 1404ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 887ms | ✓ 1112ms | ✓ 863ms | http |
| 43.161.214.161:1081 | ✓ 1348ms | ✓ 1746ms | ✓ 1346ms | ✓ 1847ms | 否 | http |
| 52.188.28.218:3128 | 否 | 否 | ✓ 226ms | ✓ 1531ms | ✓ 1086ms | http |
| 91.238.104.171:2023 | 否 | 否 | ✓ 1539ms | ✓ 1924ms | ✓ 1663ms | http |
| 165.225.122.12:11685 | ✓ 1384ms | 否 | ✓ 1408ms | 否 | ✓ 1517ms | http |
| 156.225.70.152:39151 | 否 | 否 | ✓ 1787ms | ✓ 1187ms | ✓ 867ms | http |
| 138.124.53.25:7443 | ✓ 628ms | 否 | 否 | ✓ 1601ms | ✓ 1527ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1729ms | ✓ 1850ms | ✓ 1905ms | http |
| 132.145.93.138:1080 | 否 | 否 | ✓ 1687ms | ✓ 1660ms | ✓ 1506ms | http |
| 165.225.122.16:11368 | ✓ 1099ms | 否 | ✓ 1454ms | ✓ 1488ms | ✓ 1184ms | http |
| 165.225.122.12:12302 | ✓ 1137ms | 否 | ✓ 1770ms | ✓ 1431ms | ✓ 1366ms | http |
| 165.225.122.12:11691 | ✓ 1130ms | 否 | ✓ 1781ms | ✓ 1514ms | 否 | http |
| 165.225.122.12:10919 | ✓ 1151ms | 否 | ✓ 1401ms | ✓ 1625ms | ✓ 1924ms | http |
| 165.225.122.12:12125 | ✓ 1176ms | 否 | ✓ 1370ms | ✓ 1663ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1296ms | 否 | 否 | ✓ 1877ms | ✓ 1193ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 978ms | ✓ 1191ms | ✓ 915ms | http |
| 121.230.8.136:1080 | ✓ 1113ms | 否 | ✓ 1598ms | ✓ 1706ms | 否 | http |
| 104.238.30.91:63900 | ✓ 1721ms | 否 | ✓ 1615ms | 否 | ✓ 1935ms | http |
| 121.230.8.211:1080 | ✓ 1457ms | 否 | ✓ 1962ms | ✓ 1392ms | ✓ 1099ms | http |
| 178.156.224.42:3128 | ✓ 1012ms | ✓ 1891ms | ✓ 1590ms | 否 | 否 | http |
| 34.7.88.87:3128 | ✓ 1008ms | ✓ 1844ms | ✓ 1097ms | 否 | ✓ 1639ms | http |
| 104.238.30.39:59741 | ✓ 1630ms | 否 | ✓ 1679ms | 否 | ✓ 1871ms | http |
| 44.205.216.127:80 | ✓ 149ms | ✓ 1360ms | ✓ 343ms | ✓ 1305ms | ✓ 1071ms | http |
| 172.212.68.37:3128 | ✓ 323ms | 否 | ✓ 986ms | ✓ 1051ms | ✓ 710ms | http |
| 120.55.163.237:10086 | ✓ 834ms | ✓ 1029ms | ✓ 957ms | ✓ 1055ms | ✓ 833ms | http |
| 179.96.28.58:80 | ✓ 1039ms | 否 | ✓ 948ms | 否 | ✓ 1925ms | http |
| 45.129.141.143:3128 | ✓ 1885ms | 否 | ✓ 1745ms | 否 | ✓ 1784ms | http |
| 3.214.214.245:80 | ✓ 49ms | ✓ 1390ms | ✓ 761ms | ✓ 1430ms | ✓ 1067ms | http |
| 121.40.231.103:7890 | 否 | ✓ 1002ms | ✓ 952ms | ✓ 1096ms | ✓ 817ms | http |
| 101.47.73.135:3128 | ✓ 1561ms | 否 | ✓ 1123ms | ✓ 1510ms | 否 | http |
| 195.123.209.48:3128 | ✓ 943ms | 否 | ✓ 1690ms | ✓ 1666ms | ✓ 1534ms | http |
| 223.16.170.103:80 | ✓ 1110ms | 否 | ✓ 1541ms | ✓ 1530ms | ✓ 1395ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1455ms | ✓ 1380ms | ✓ 1083ms | http |
| 36.147.78.166:443 | 否 | ✓ 1893ms | ✓ 1904ms | 否 | ✓ 1889ms | http |
| 223.16.170.103:3128 | ✓ 1793ms | 否 | ✓ 1259ms | ✓ 1304ms | ✓ 1321ms | http |
| 14.143.222.113:10158 | ✓ 959ms | 否 | ✓ 1760ms | ✓ 1430ms | 否 | http |
| 61.72.110.94:3128 | 否 | ✓ 1328ms | ✓ 1062ms | ✓ 1492ms | ✓ 1991ms | http |
| 38.180.2.107:3128 | ✓ 1938ms | ✓ 1959ms | 否 | ✓ 1904ms | ✓ 1784ms | http |
| 37.27.100.79:443 | ✓ 667ms | 否 | ✓ 1812ms | ✓ 1544ms | 否 | http |
| 165.225.122.12:11962 | 否 | 否 | ✓ 1442ms | ✓ 1623ms | ✓ 1345ms | http |
| 45.136.198.40:3128 | 否 | ✓ 1761ms | ✓ 1902ms | ✓ 1811ms | 否 | http |
| 103.215.36.88:16894 | ✓ 1775ms | ✓ 1320ms | ✓ 1829ms | ✓ 1375ms | ✓ 1927ms | http |
| 158.160.215.167:8127 | ✓ 1244ms | ✓ 1770ms | ✓ 1711ms | 否 | 否 | http |

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
