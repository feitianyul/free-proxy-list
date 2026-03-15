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

最后更新：2026-03-15 00:23:08 UTC（2026-03-15 08:23:08 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.145.203.135:8443 | ✓ 292ms | ✓ 909ms | ✓ 328ms | ✓ 939ms | ✓ 733ms | http |
| 38.145.218.82:8443 | ✓ 387ms | ✓ 1675ms | ✓ 326ms | ✓ 968ms | ✓ 700ms | http |
| 205.209.118.30:3138 | ✓ 653ms | ✓ 907ms | ✓ 731ms | ✓ 1055ms | ✓ 809ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 910ms | ✓ 1211ms | ✓ 859ms | http |
| 113.160.132.26:8080 | ✓ 1072ms | ✓ 1515ms | ✓ 1238ms | ✓ 1760ms | ✓ 1907ms | http |
| 45.167.124.52:8080 | ✓ 1139ms | ✓ 1537ms | ✓ 544ms | ✓ 1584ms | ✓ 1305ms | http |
| 47.101.149.27:9010 | ✓ 1565ms | ✓ 1588ms | ✓ 1570ms | 否 | ✓ 1602ms | http |
| 138.124.53.25:7443 | ✓ 460ms | ✓ 1785ms | ✓ 1426ms | 否 | 否 | http |
| 95.3.9.78:3128 | ✓ 1429ms | 否 | 否 | ✓ 1860ms | ✓ 1394ms | http |
| 35.225.22.61:80 | ✓ 677ms | 否 | ✓ 1076ms | ✓ 1153ms | 否 | http |
| 95.3.9.78:8080 | ✓ 1418ms | ✓ 1928ms | ✓ 1942ms | ✓ 1760ms | ✓ 1223ms | http |
| 165.227.5.10:8888 | ✓ 368ms | ✓ 895ms | ✓ 783ms | ✓ 1270ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1118ms | ✓ 1388ms | ✓ 1014ms | ✓ 1622ms | ✓ 999ms | http |
| 81.70.169.194:80 | ✓ 1207ms | ✓ 1511ms | ✓ 1238ms | ✓ 1409ms | ✓ 1168ms | http |
| 150.230.249.50:1080 | 否 | ✓ 1975ms | ✓ 1151ms | ✓ 1299ms | ✓ 1002ms | http |
| 62.60.177.204:34094 | ✓ 1028ms | ✓ 1235ms | ✓ 971ms | ✓ 1023ms | ✓ 958ms | http |
| 86.53.183.16:1080 | ✓ 426ms | ✓ 1727ms | 否 | ✓ 1938ms | ✓ 1361ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1387ms | ✓ 1506ms | ✓ 1369ms | http |
| 190.12.150.244:999 | ✓ 1491ms | ✓ 1797ms | ✓ 904ms | ✓ 1616ms | ✓ 1365ms | http |
| 43.167.227.161:1080 | 否 | ✓ 1578ms | ✓ 645ms | ✓ 1095ms | ✓ 817ms | http |
| 85.198.96.242:3128 | ✓ 983ms | 否 | 否 | ✓ 1584ms | ✓ 1206ms | http |
| 144.31.25.69:21064 | ✓ 1848ms | 否 | ✓ 944ms | 否 | ✓ 1928ms | http |
| 14.225.212.37:7890 | ✓ 1647ms | ✓ 1737ms | 否 | ✓ 1336ms | 否 | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 67ms | ✓ 1062ms | ✓ 769ms | http |
| 120.92.212.16:8890 | ✓ 1170ms | 否 | ✓ 1164ms | ✓ 1428ms | 否 | http |
| 185.241.5.57:3128 | ✓ 1164ms | 否 | ✓ 1851ms | 否 | ✓ 1500ms | http |
| 121.230.8.109:1080 | ✓ 1231ms | ✓ 1669ms | ✓ 1440ms | ✓ 1875ms | ✓ 1379ms | http |
| 121.230.8.97:1080 | ✓ 1626ms | ✓ 1600ms | ✓ 1314ms | ✓ 1946ms | ✓ 1249ms | http |
| 59.46.216.131:30001 | ✓ 1125ms | ✓ 1648ms | ✓ 1246ms | ✓ 1595ms | ✓ 1314ms | http |
| 121.40.231.103:7890 | ✓ 1044ms | ✓ 1255ms | ✓ 1066ms | ✓ 1296ms | ✓ 1061ms | http |
| 91.233.223.147:3128 | ✓ 1009ms | 否 | ✓ 1007ms | 否 | ✓ 1894ms | http |
| 120.92.212.16:7890 | ✓ 1155ms | 否 | ✓ 1969ms | ✓ 1433ms | ✓ 1155ms | http |
| 113.255.59.226:8080 | ✓ 1730ms | 否 | 否 | ✓ 1697ms | ✓ 1362ms | http |
| 106.117.208.101:7890 | ✓ 1225ms | ✓ 1442ms | ✓ 1222ms | ✓ 1527ms | ✓ 1439ms | http |
| 101.43.255.96:80 | ✓ 1256ms | ✓ 1492ms | ✓ 1062ms | ✓ 1421ms | ✓ 1166ms | http |
| 195.158.8.123:3128 | ✓ 1570ms | 否 | ✓ 1698ms | 否 | ✓ 1925ms | http |
| 192.73.243.98:3128 | ✓ 250ms | ✓ 1526ms | ✓ 1829ms | ✓ 1098ms | ✓ 728ms | http |
| 194.147.90.23:3128 | ✓ 829ms | 否 | ✓ 946ms | ✓ 1808ms | ✓ 1223ms | http |
| 168.138.202.218:3128 | ✓ 1856ms | 否 | ✓ 1749ms | ✓ 1185ms | ✓ 928ms | http |
| 45.77.246.231:80 | ✓ 1613ms | 否 | ✓ 1691ms | ✓ 1407ms | ✓ 1143ms | http |
| 88.80.150.82:8080 | ✓ 1158ms | 否 | 否 | ✓ 1854ms | ✓ 1626ms | https |
| 164.92.148.68:3128 | ✓ 999ms | ✓ 1996ms | ✓ 1813ms | ✓ 1505ms | ✓ 1570ms | http |
| 45.129.141.143:3128 | ✓ 1090ms | 否 | ✓ 1763ms | ✓ 1924ms | 否 | http |
| 150.249.255.91:3128 | ✓ 717ms | ✓ 1177ms | ✓ 827ms | 否 | ✓ 1552ms | http |
| 62.234.206.73:3128 | ✓ 1179ms | ✓ 1507ms | ✓ 1133ms | ✓ 1528ms | ✓ 1150ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1274ms | ✓ 1992ms | 否 | ✓ 1338ms | http |
| 45.136.198.40:3128 | ✓ 978ms | ✓ 1406ms | ✓ 1168ms | ✓ 1775ms | ✓ 1744ms | http |
| 62.113.119.14:8080 | ✓ 869ms | ✓ 1564ms | ✓ 1589ms | 否 | 否 | http |
| 38.180.2.107:3128 | ✓ 878ms | 否 | ✓ 1706ms | ✓ 1976ms | ✓ 1617ms | http |
| 113.176.92.71:3128 | ✓ 1801ms | ✓ 1545ms | ✓ 1482ms | ✓ 1713ms | ✓ 1420ms | http |
| 120.55.163.237:10086 | ✓ 1794ms | ✓ 1573ms | 否 | 否 | ✓ 1680ms | http |
| 1.225.116.115:1080 | ✓ 1663ms | 否 | 否 | ✓ 1487ms | ✓ 1642ms | http |
| 194.5.212.40:8080 | ✓ 1442ms | ✓ 1319ms | ✓ 1364ms | 否 | ✓ 1640ms | http |
| 38.145.208.37:8443 | ✓ 591ms | ✓ 921ms | ✓ 792ms | ✓ 1218ms | ✓ 734ms | http |
| 38.145.208.148:8443 | 否 | ✓ 1761ms | ✓ 318ms | ✓ 999ms | ✓ 750ms | http |
| 45.136.130.228:8443 | 否 | ✓ 1761ms | ✓ 317ms | ✓ 1000ms | ✓ 760ms | http |
| 38.145.208.144:8443 | 否 | ✓ 1775ms | ✓ 323ms | ✓ 972ms | ✓ 779ms | http |
| 1.234.153.14:80 | ✓ 1995ms | ✓ 1447ms | ✓ 1258ms | ✓ 1126ms | ✓ 884ms | http |
| 61.52.131.172:8443 | ✓ 882ms | ✓ 1198ms | ✓ 917ms | ✓ 1217ms | ✓ 940ms | http |
| 65.108.203.37:18080 | ✓ 1441ms | ✓ 1953ms | ✓ 1602ms | 否 | 否 | http |
| 65.108.203.35:18080 | ✓ 1214ms | ✓ 1697ms | 否 | ✓ 1892ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1584ms | 否 | 否 | ✓ 1701ms | ✓ 1593ms | http |
| 38.145.208.138:8447 | ✓ 616ms | ✓ 963ms | ✓ 322ms | ✓ 1593ms | ✓ 706ms | http |
| 38.145.208.135:8443 | 否 | ✓ 1092ms | ✓ 507ms | ✓ 1568ms | ✓ 711ms | http |
| 38.145.208.137:8443 | ✓ 366ms | ✓ 1360ms | ✓ 304ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 270ms | 否 | 否 | ✓ 1078ms | ✓ 815ms | http |
| 45.140.147.82:1081 | ✓ 950ms | ✓ 1563ms | ✓ 1469ms | ✓ 1442ms | ✓ 870ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1570ms | ✓ 1327ms | ✓ 1387ms | http |

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
