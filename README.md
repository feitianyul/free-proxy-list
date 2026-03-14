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

最后更新：2026-03-14 21:26:55 UTC（2026-03-15 05:26:55 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 109.73.195.10:8888 | ✓ 1164ms | 否 | ✓ 1513ms | ✓ 1708ms | 否 | http |
| 62.60.177.204:34094 | 否 | ✓ 1383ms | ✓ 1464ms | ✓ 1053ms | ✓ 740ms | http |
| 95.3.9.78:8080 | ✓ 1721ms | ✓ 1980ms | 否 | ✓ 1713ms | ✓ 1349ms | http |
| 116.80.65.79:3172 | 否 | 否 | ✓ 1697ms | ✓ 1961ms | ✓ 1815ms | http |
| 116.80.96.107:3172 | 否 | 否 | ✓ 1697ms | ✓ 1987ms | ✓ 1784ms | http |
| 35.225.22.61:80 | ✓ 663ms | ✓ 1036ms | ✓ 726ms | 否 | ✓ 932ms | http |
| 38.145.203.135:8443 | ✓ 434ms | ✓ 960ms | ✓ 536ms | ✓ 978ms | ✓ 797ms | http |
| 103.84.95.54:7890 | ✓ 752ms | 否 | ✓ 790ms | ✓ 1292ms | 否 | http |
| 101.43.127.100:8877 | ✓ 937ms | ✓ 1262ms | ✓ 943ms | ✓ 1310ms | ✓ 939ms | http |
| 81.70.169.194:80 | ✓ 1155ms | ✓ 1387ms | ✓ 1038ms | ✓ 1372ms | ✓ 1070ms | http |
| 101.43.255.96:80 | ✓ 1069ms | ✓ 1602ms | ✓ 1104ms | ✓ 1363ms | ✓ 1157ms | http |
| 2.56.122.146:10808 | ✓ 1117ms | ✓ 1306ms | ✓ 917ms | ✓ 1403ms | ✓ 1878ms | http |
| 150.230.249.50:1080 | ✓ 1142ms | ✓ 1693ms | 否 | 否 | ✓ 917ms | http |
| 85.198.96.242:3128 | ✓ 572ms | 否 | ✓ 941ms | ✓ 1622ms | ✓ 1301ms | http |
| 95.3.9.78:3128 | ✓ 778ms | ✓ 1753ms | ✓ 1637ms | 否 | 否 | http |
| 45.136.130.223:8443 | ✓ 1067ms | ✓ 912ms | ✓ 540ms | ✓ 964ms | ✓ 753ms | http |
| 45.136.131.42:8447 | 否 | ✓ 1131ms | ✓ 1066ms | 否 | ✓ 1415ms | http |
| 59.46.216.131:30001 | ✓ 1055ms | ✓ 1432ms | ✓ 1261ms | 否 | 否 | http |
| 185.115.74.185:8080 | ✓ 1095ms | ✓ 1593ms | ✓ 1697ms | 否 | 否 | http |
| 38.145.218.82:8443 | ✓ 869ms | ✓ 987ms | ✓ 925ms | ✓ 1572ms | ✓ 1276ms | http |
| 165.227.5.10:8888 | ✓ 846ms | 否 | 否 | ✓ 1900ms | ✓ 1355ms | http |
| 193.148.58.165:3128 | ✓ 709ms | 否 | ✓ 1760ms | ✓ 1894ms | ✓ 1845ms | http |
| 47.101.159.19:8899 | 否 | 否 | ✓ 1850ms | ✓ 1297ms | ✓ 981ms | http |
| 113.160.132.26:8080 | ✓ 1898ms | ✓ 1838ms | ✓ 1577ms | 否 | 否 | http |
| 117.86.6.201:1080 | ✓ 1356ms | ✓ 1450ms | ✓ 1120ms | ✓ 1534ms | ✓ 1105ms | http |
| 205.209.118.30:3138 | ✓ 1212ms | ✓ 1822ms | 否 | ✓ 1211ms | 否 | http |
| 14.225.212.37:7890 | ✓ 1406ms | 否 | 否 | ✓ 1217ms | ✓ 1042ms | http |
| 45.136.131.32:8443 | ✓ 347ms | ✓ 924ms | ✓ 339ms | ✓ 940ms | ✓ 764ms | http |
| 38.145.208.37:8443 | ✓ 360ms | ✓ 927ms | ✓ 330ms | ✓ 963ms | ✓ 726ms | http |
| 45.136.131.33:8443 | ✓ 323ms | ✓ 925ms | ✓ 350ms | ✓ 996ms | ✓ 770ms | http |
| 38.145.208.136:8443 | ✓ 345ms | ✓ 923ms | ✓ 337ms | ✓ 997ms | ✓ 768ms | http |
| 38.145.208.132:8443 | ✓ 349ms | ✓ 932ms | ✓ 341ms | ✓ 963ms | ✓ 759ms | http |
| 45.136.131.30:8447 | ✓ 335ms | ✓ 961ms | ✓ 378ms | ✓ 957ms | ✓ 752ms | http |
| 45.136.130.230:8443 | ✓ 368ms | ✓ 921ms | ✓ 375ms | ✓ 961ms | ✓ 750ms | http |
| 45.136.131.35:8443 | ✓ 357ms | ✓ 908ms | ✓ 323ms | ✓ 1003ms | ✓ 787ms | http |
| 38.145.208.138:8447 | ✓ 405ms | ✓ 954ms | ✓ 363ms | ✓ 989ms | ✓ 731ms | http |
| 45.136.131.31:8443 | ✓ 370ms | ✓ 953ms | ✓ 327ms | ✓ 1000ms | ✓ 733ms | http |
| 38.145.208.96:8443 | ✓ 379ms | ✓ 908ms | ✓ 346ms | ✓ 1040ms | ✓ 745ms | http |
| 38.145.203.161:8443 | ✓ 403ms | ✓ 938ms | ✓ 338ms | ✓ 984ms | ✓ 811ms | http |
| 38.145.203.246:8443 | ✓ 418ms | ✓ 933ms | ✓ 376ms | ✓ 1003ms | ✓ 761ms | http |
| 38.145.208.131:8443 | ✓ 341ms | ✓ 967ms | ✓ 335ms | ✓ 1005ms | ✓ 791ms | http |
| 38.145.208.149:8443 | ✓ 410ms | ✓ 923ms | ✓ 350ms | ✓ 1015ms | ✓ 789ms | http |
| 38.145.208.98:8443 | ✓ 389ms | ✓ 907ms | ✓ 347ms | ✓ 1056ms | ✓ 803ms | http |
| 38.145.208.148:8443 | ✓ 411ms | ✓ 904ms | ✓ 328ms | ✓ 1062ms | ✓ 746ms | http |
| 38.145.208.97:8443 | ✓ 445ms | ✓ 932ms | ✓ 334ms | ✓ 1018ms | ✓ 751ms | http |
| 38.145.208.93:8443 | ✓ 434ms | ✓ 1012ms | ✓ 377ms | ✓ 1015ms | ✓ 751ms | http |
| 38.145.203.162:8443 | ✓ 432ms | ✓ 1470ms | ✓ 337ms | ✓ 945ms | ✓ 742ms | http |
| 38.145.208.94:8443 | ✓ 405ms | ✓ 1476ms | ✓ 349ms | ✓ 967ms | ✓ 763ms | http |
| 38.145.203.245:8443 | ✓ 385ms | ✓ 1508ms | ✓ 365ms | ✓ 1015ms | ✓ 754ms | http |
| 45.136.130.162:8443 | ✓ 509ms | ✓ 1260ms | ✓ 1013ms | ✓ 1048ms | ✓ 756ms | http |
| 45.136.130.156:8443 | ✓ 534ms | ✓ 1265ms | ✓ 964ms | ✓ 1118ms | ✓ 757ms | http |
| 45.136.130.155:8443 | ✓ 522ms | ✓ 1265ms | ✓ 966ms | ✓ 1124ms | ✓ 786ms | http |
| 168.235.110.63:3128 | ✓ 379ms | 否 | ✓ 910ms | ✓ 1558ms | ✓ 764ms | http |
| 138.124.53.25:7443 | ✓ 1453ms | 否 | ✓ 1233ms | ✓ 1970ms | 否 | http |
| 45.136.131.39:8443 | ✓ 1022ms | ✓ 905ms | ✓ 1338ms | ✓ 1056ms | ✓ 889ms | http |
| 103.183.10.169:3125 | ✓ 1728ms | 否 | ✓ 1440ms | ✓ 1633ms | ✓ 1677ms | http |
| 121.230.8.34:1080 | ✓ 1078ms | ✓ 1417ms | ✓ 1311ms | ✓ 1546ms | ✓ 1326ms | http |
| 192.71.213.85:9812 | ✓ 1293ms | 否 | ✓ 1319ms | ✓ 1716ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1170ms | ✓ 1884ms | 否 | 否 | ✓ 1934ms | https |
| 165.225.113.220:10958 | 否 | 否 | ✓ 921ms | ✓ 1207ms | ✓ 964ms | http |
| 194.5.212.40:8080 | ✓ 1228ms | ✓ 1824ms | ✓ 1516ms | ✓ 1992ms | ✓ 1527ms | http |
| 103.82.23.118:5221 | ✓ 1845ms | 否 | ✓ 1829ms | ✓ 1651ms | ✓ 1904ms | http |
| 116.80.65.75:3172 | ✓ 1635ms | 否 | 否 | ✓ 1971ms | ✓ 1809ms | http |
| 120.92.212.16:8890 | ✓ 1082ms | ✓ 1377ms | 否 | 否 | ✓ 1067ms | http |
| 120.92.212.16:7890 | ✓ 1096ms | ✓ 1414ms | 否 | ✓ 1380ms | 否 | http |
| 38.145.218.101:8447 | 否 | ✓ 993ms | ✓ 840ms | ✓ 968ms | ✓ 759ms | http |
| 45.136.131.28:8447 | 否 | ✓ 964ms | ✓ 1032ms | ✓ 1166ms | ✓ 772ms | http |
| 106.117.208.101:7890 | ✓ 1118ms | ✓ 1436ms | 否 | ✓ 1771ms | 否 | http |
| 150.249.255.91:3128 | ✓ 1543ms | ✓ 1933ms | ✓ 749ms | ✓ 1043ms | ✓ 835ms | http |
| 47.252.41.213:443 | ✓ 230ms | ✓ 1100ms | ✓ 1072ms | ✓ 1211ms | ✓ 1216ms | http |
| 178.156.224.42:3128 | ✓ 1042ms | 否 | ✓ 1596ms | 否 | ✓ 1742ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1054ms | ✓ 1210ms | ✓ 947ms | http |
| 210.77.29.245:7890 | ✓ 1655ms | 否 | ✓ 1272ms | ✓ 1515ms | ✓ 1739ms | http |
| 116.80.96.102:3172 | ✓ 1663ms | 否 | 否 | ✓ 1967ms | ✓ 1768ms | http |
| 103.184.99.194:8080 | 否 | 否 | ✓ 1952ms | ✓ 1602ms | ✓ 1608ms | http |
| 89.185.85.138:1080 | ✓ 1170ms | 否 | 否 | ✓ 1773ms | ✓ 1357ms | http |
| 103.39.51.190:8080 | ✓ 1943ms | 否 | 否 | ✓ 1979ms | ✓ 1934ms | http |
| 38.145.208.187:8443 | ✓ 363ms | ✓ 955ms | ✓ 572ms | ✓ 995ms | ✓ 757ms | http |
| 45.136.198.40:3128 | ✓ 1099ms | 否 | ✓ 1434ms | 否 | ✓ 1677ms | http |
| 101.47.73.135:3128 | ✓ 1133ms | 否 | ✓ 1179ms | ✓ 1299ms | ✓ 1364ms | http |
| 86.53.183.16:1080 | 否 | 否 | ✓ 1100ms | ✓ 1376ms | ✓ 1303ms | http |
| 61.52.131.172:8443 | ✓ 1030ms | ✓ 1263ms | ✓ 988ms | ✓ 1307ms | ✓ 1054ms | http |
| 103.82.23.118:5195 | ✓ 1855ms | 否 | ✓ 1357ms | ✓ 1690ms | 否 | http |
| 45.140.147.155:1081 | ✓ 1024ms | 否 | ✓ 687ms | ✓ 1799ms | ✓ 1185ms | http |
| 193.23.200.251:10808 | ✓ 971ms | ✓ 1835ms | ✓ 1970ms | 否 | ✓ 1749ms | http |

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
