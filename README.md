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

最后更新：2026-05-07 07:29:23 UTC（2026-05-07 15:29:23 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 193.160.209.58:1080 | ✓ 1324ms | 否 | ✓ 1985ms | 否 | ✓ 1787ms | http |
| 181.119.97.24:999 | 否 | 否 | ✓ 1615ms | ✓ 1745ms | ✓ 1490ms | http |
| 47.77.216.82:1080 | ✓ 394ms | 否 | ✓ 429ms | ✓ 1056ms | ✓ 1116ms | http |
| 62.113.119.14:8080 | ✓ 523ms | ✓ 1365ms | ✓ 907ms | ✓ 1406ms | ✓ 1031ms | http |
| 1.231.81.166:3128 | ✓ 1320ms | ✓ 1253ms | ✓ 1076ms | ✓ 1230ms | ✓ 1098ms | http |
| 45.125.67.37:8443 | ✓ 1245ms | 否 | ✓ 910ms | ✓ 1330ms | ✓ 1329ms | http |
| 84.47.150.125:1080 | ✓ 1342ms | 否 | 否 | ✓ 1863ms | ✓ 1590ms | http |
| 118.113.244.152:1080 | ✓ 1623ms | ✓ 1905ms | ✓ 1450ms | ✓ 1867ms | ✓ 1461ms | http |
| 113.160.132.26:8080 | ✓ 1766ms | ✓ 1898ms | ✓ 1262ms | 否 | ✓ 1616ms | http |
| 62.133.60.126:24558 | ✓ 521ms | 否 | 否 | ✓ 1762ms | ✓ 1510ms | http |
| 158.160.215.167:8127 | ✓ 912ms | 否 | ✓ 1587ms | 否 | ✓ 1603ms | http |
| 8.219.97.248:80 | ✓ 1518ms | 否 | ✓ 1793ms | 否 | ✓ 1551ms | http |
| 5.161.50.82:8118 | ✓ 1568ms | 否 | ✓ 662ms | 否 | ✓ 1688ms | http |
| 91.242.229.129:8092 | ✓ 1301ms | 否 | 否 | ✓ 1521ms | ✓ 1166ms | http |
| 218.108.131.186:17890 | 否 | 否 | ✓ 1173ms | ✓ 1323ms | ✓ 1342ms | http |
| 20.164.75.153:8080 | ✓ 1157ms | 否 | ✓ 1651ms | 否 | ✓ 1801ms | http |
| 77.110.107.80:8080 | ✓ 921ms | 否 | ✓ 1011ms | ✓ 1973ms | ✓ 1197ms | http |
| 45.59.122.132:80 | ✓ 1436ms | 否 | ✓ 782ms | ✓ 1283ms | ✓ 1388ms | http |
| 43.156.132.113:3128 | ✓ 1373ms | 否 | ✓ 955ms | ✓ 1661ms | ✓ 1003ms | http |
| 43.133.44.89:8888 | ✓ 1929ms | 否 | 否 | ✓ 1248ms | ✓ 1979ms | http |
| 45.236.129.64:3128 | ✓ 643ms | ✓ 1682ms | ✓ 481ms | ✓ 1903ms | ✓ 1448ms | http |
| 5.252.33.13:2025 | ✓ 1291ms | 否 | ✓ 1127ms | ✓ 1974ms | ✓ 1741ms | http |
| 59.46.216.131:30001 | ✓ 1207ms | ✓ 1546ms | 否 | 否 | ✓ 1313ms | http |
| 38.211.245.34:999 | ✓ 1023ms | 否 | ✓ 969ms | 否 | ✓ 1854ms | http |
| 152.70.91.193:40000 | ✓ 1928ms | 否 | 否 | ✓ 1716ms | ✓ 1448ms | http |
| 20.27.11.248:8561 | ✓ 1487ms | ✓ 1059ms | ✓ 748ms | ✓ 1081ms | ✓ 1753ms | http |
| 20.27.14.220:8561 | ✓ 1496ms | ✓ 1094ms | ✓ 713ms | ✓ 1080ms | ✓ 1767ms | http |
| 20.27.13.35:8561 | 否 | ✓ 1137ms | ✓ 655ms | ✓ 1439ms | 否 | http |
| 77.110.107.80:1080 | ✓ 1047ms | ✓ 1573ms | ✓ 1696ms | 否 | ✓ 966ms | http |
| 91.233.223.147:3128 | ✓ 835ms | 否 | ✓ 937ms | ✓ 1995ms | ✓ 1625ms | http |
| 157.0.142.246:10057 | ✓ 1233ms | ✓ 1621ms | ✓ 1193ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 985ms | 否 | ✓ 1393ms | 否 | ✓ 1790ms | http |
| 157.230.220.25:4857 | ✓ 230ms | ✓ 1016ms | ✓ 908ms | ✓ 889ms | 否 | http |
| 64.188.77.221:3128 | ✓ 1887ms | ✓ 1801ms | ✓ 744ms | ✓ 1572ms | ✓ 1348ms | http |
| 201.144.20.238:3128 | ✓ 746ms | ✓ 1298ms | ✓ 1271ms | ✓ 1409ms | 否 | http |
| 77.93.89.128:47146 | ✓ 1327ms | 否 | ✓ 1197ms | ✓ 1325ms | ✓ 1036ms | http |
| 152.32.132.190:7890 | ✓ 1039ms | 否 | ✓ 1384ms | 否 | ✓ 915ms | http |
| 206.206.126.177:2412 | ✓ 1567ms | 否 | ✓ 1257ms | 否 | ✓ 1121ms | http |
| 64.188.77.26:3128 | ✓ 996ms | ✓ 1945ms | 否 | ✓ 1372ms | ✓ 1175ms | http |
| 223.16.170.103:80 | ✓ 1191ms | 否 | 否 | ✓ 1303ms | ✓ 1819ms | http |
| 77.110.119.136:3128 | ✓ 597ms | 否 | ✓ 236ms | ✓ 932ms | 否 | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1332ms | ✓ 1297ms | ✓ 1220ms | http |
| 8.154.21.175:3128 | 否 | ✓ 1350ms | ✓ 1055ms | ✓ 1341ms | ✓ 1126ms | http |
| 38.194.254.134:999 | ✓ 937ms | ✓ 1582ms | ✓ 1363ms | 否 | 否 | http |
| 190.52.108.145:999 | 否 | ✓ 1438ms | ✓ 1166ms | ✓ 1543ms | 否 | http |
| 116.118.48.147:3128 | ✓ 1561ms | 否 | ✓ 1680ms | ✓ 1396ms | ✓ 1250ms | http |
| 47.112.25.109:7890 | ✓ 1171ms | 否 | ✓ 1937ms | 否 | ✓ 1521ms | http |
| 103.250.128.8:8082 | ✓ 1528ms | 否 | ✓ 1511ms | ✓ 1611ms | ✓ 1754ms | http |
| 144.124.227.88:3128 | ✓ 1436ms | 否 | ✓ 1341ms | ✓ 1745ms | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1186ms | ✓ 806ms | ✓ 1062ms | ✓ 804ms | http |
| 185.191.236.162:3128 | ✓ 1627ms | ✓ 1547ms | ✓ 825ms | ✓ 1540ms | ✓ 1045ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1544ms | ✓ 1542ms | ✓ 1326ms | http |
| 34.96.238.40:8080 | ✓ 1972ms | 否 | 否 | ✓ 1507ms | ✓ 1291ms | http |
| 47.79.39.142:30000 | 否 | 否 | ✓ 1693ms | ✓ 1578ms | ✓ 1834ms | http |
| 160.238.65.5:3128 | ✓ 1158ms | 否 | ✓ 1724ms | 否 | ✓ 1363ms | http |
| 160.238.65.9:3128 | ✓ 1155ms | 否 | ✓ 1717ms | 否 | ✓ 1434ms | http |
| 160.238.65.8:3128 | ✓ 1282ms | 否 | ✓ 828ms | ✓ 1301ms | ✓ 1329ms | http |
| 160.238.65.2:3128 | ✓ 1174ms | 否 | ✓ 1718ms | 否 | ✓ 1641ms | http |
| 91.107.124.215:3128 | ✓ 1204ms | 否 | ✓ 847ms | 否 | ✓ 1802ms | http |
| 160.238.65.3:3128 | ✓ 1158ms | 否 | ✓ 1718ms | 否 | ✓ 1414ms | http |
| 160.238.65.4:3128 | ✓ 1843ms | ✓ 1442ms | 否 | 否 | ✓ 1001ms | http |
| 178.63.155.151:8978 | ✓ 1220ms | ✓ 1482ms | ✓ 926ms | ✓ 1662ms | ✓ 1413ms | http |
| 112.65.132.181:3128 | ✓ 1345ms | ✓ 1212ms | ✓ 992ms | ✓ 1271ms | ✓ 1056ms | http |
| 38.188.247.12:999 | ✓ 1051ms | ✓ 1351ms | ✓ 1078ms | ✓ 1282ms | ✓ 1021ms | http |
| 177.234.221.20:999 | ✓ 1256ms | ✓ 1795ms | ✓ 1036ms | 否 | 否 | http |
| 2.27.32.81:3128 | ✓ 1860ms | ✓ 1780ms | ✓ 835ms | ✓ 1953ms | 否 | http |
| 109.234.38.35:3128 | ✓ 1960ms | ✓ 1532ms | ✓ 1387ms | ✓ 1872ms | ✓ 1708ms | http |
| 3.101.133.120:80 | ✓ 1258ms | ✓ 1949ms | ✓ 631ms | ✓ 1578ms | ✓ 1202ms | http |
| 160.238.65.6:3128 | ✓ 1647ms | 否 | ✓ 1813ms | ✓ 1960ms | 否 | http |
| 200.110.173.240:999 | ✓ 1084ms | 否 | 否 | ✓ 1896ms | ✓ 1662ms | http |
| 194.59.247.34:10808 | ✓ 1487ms | ✓ 1401ms | ✓ 1015ms | ✓ 1858ms | ✓ 1392ms | http |
| 103.35.190.69:1082 | ✓ 215ms | ✓ 1066ms | ✓ 103ms | ✓ 1216ms | 否 | http |
| 130.61.120.40:3128 | ✓ 848ms | 否 | 否 | ✓ 1649ms | ✓ 1161ms | http |
| 103.106.219.107:8081 | ✓ 1722ms | 否 | ✓ 1692ms | ✓ 1807ms | ✓ 1945ms | http |
| 61.52.131.172:8443 | ✓ 1098ms | ✓ 1424ms | ✓ 1059ms | ✓ 1441ms | ✓ 1157ms | http |
| 13.60.181.61:12484 | ✓ 640ms | 否 | ✓ 1563ms | ✓ 1945ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1942ms | 否 | ✓ 1174ms | ✓ 1661ms | ✓ 1401ms | http |

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
