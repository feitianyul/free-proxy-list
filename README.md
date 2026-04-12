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

最后更新：2026-04-12 15:33:48 UTC（2026-04-12 23:33:48 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 548ms | 否 | 否 | ✓ 1196ms | ✓ 975ms | http |
| 147.161.210.140:8800 | ✓ 1770ms | 否 | ✓ 1235ms | 否 | ✓ 1205ms | http |
| 167.103.115.102:8800 | ✓ 1743ms | 否 | ✓ 1283ms | ✓ 1283ms | ✓ 1292ms | http |
| 113.160.132.26:8080 | ✓ 1617ms | 否 | ✓ 1566ms | ✓ 1814ms | ✓ 1088ms | http |
| 167.103.34.108:8800 | ✓ 1823ms | 否 | ✓ 1531ms | ✓ 1511ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1811ms | 否 | 否 | ✓ 1409ms | ✓ 1135ms | http |
| 20.210.39.153:8561 | ✓ 1772ms | ✓ 1490ms | ✓ 1512ms | ✓ 1678ms | ✓ 1543ms | http |
| 20.78.118.91:8561 | ✓ 1772ms | ✓ 1885ms | ✓ 1356ms | ✓ 1639ms | ✓ 1421ms | http |
| 20.27.15.49:8561 | 否 | 否 | ✓ 1561ms | ✓ 1963ms | ✓ 1627ms | http |
| 45.167.124.52:8080 | ✓ 1080ms | ✓ 1591ms | 否 | ✓ 1677ms | ✓ 1737ms | http |
| 95.214.9.93:3128 | ✓ 963ms | 否 | ✓ 650ms | 否 | ✓ 1705ms | http |
| 79.132.136.58:3128 | ✓ 702ms | 否 | ✓ 1695ms | ✓ 1825ms | ✓ 1535ms | http |
| 223.84.151.86:30005 | ✓ 1468ms | ✓ 1268ms | ✓ 1350ms | ✓ 1899ms | ✓ 1619ms | http |
| 5.104.87.17:8051 | ✓ 1344ms | 否 | ✓ 1198ms | ✓ 1475ms | ✓ 1169ms | http |
| 167.103.144.127:8800 | ✓ 1877ms | 否 | ✓ 1359ms | ✓ 1752ms | ✓ 1530ms | http |
| 59.46.216.131:30001 | ✓ 1278ms | ✓ 1629ms | ✓ 1375ms | 否 | 否 | http |
| 45.167.125.21:999 | ✓ 658ms | ✓ 1931ms | 否 | ✓ 1757ms | ✓ 1414ms | http |
| 103.125.181.135:9999 | ✓ 1052ms | 否 | ✓ 1708ms | ✓ 1635ms | ✓ 1344ms | http |
| 20.210.76.104:8561 | ✓ 1127ms | 否 | ✓ 1223ms | ✓ 1632ms | ✓ 1204ms | http |
| 20.210.76.175:8561 | ✓ 1156ms | 否 | ✓ 1237ms | ✓ 1600ms | ✓ 1197ms | http |
| 218.108.131.186:17890 | ✓ 995ms | ✓ 1258ms | ✓ 1251ms | ✓ 1271ms | ✓ 1042ms | http |
| 139.159.99.242:8080 | 否 | ✓ 1203ms | ✓ 1014ms | 否 | ✓ 1035ms | http |
| 167.103.31.122:8800 | ✓ 1368ms | 否 | ✓ 1533ms | 否 | ✓ 1448ms | http |
| 205.164.46.6:3094 | ✓ 330ms | ✓ 937ms | 否 | ✓ 1331ms | 否 | http |
| 20.78.26.206:8561 | ✓ 630ms | ✓ 1129ms | ✓ 667ms | ✓ 968ms | ✓ 756ms | http |
| 147.161.239.240:8800 | ✓ 596ms | ✓ 1503ms | ✓ 1285ms | ✓ 1547ms | ✓ 1267ms | http |
| 8.219.195.129:1080 | ✓ 1662ms | ✓ 1921ms | ✓ 1053ms | ✓ 1246ms | ✓ 1032ms | http |
| 159.223.225.118:8888 | ✓ 401ms | 否 | ✓ 1512ms | 否 | ✓ 1222ms | http |
| 91.233.223.147:3128 | ✓ 762ms | 否 | ✓ 1165ms | ✓ 1895ms | 否 | http |
| 177.234.217.88:999 | ✓ 1372ms | ✓ 1987ms | 否 | 否 | ✓ 1457ms | http |
| 20.27.13.35:8561 | ✓ 681ms | ✓ 1050ms | ✓ 890ms | ✓ 1044ms | ✓ 768ms | http |
| 20.27.14.220:8561 | ✓ 681ms | ✓ 1380ms | ✓ 715ms | ✓ 1040ms | ✓ 869ms | http |
| 20.27.15.111:8561 | ✓ 681ms | ✓ 1637ms | ✓ 700ms | ✓ 1043ms | ✓ 830ms | http |
| 20.27.11.248:8561 | ✓ 647ms | ✓ 1834ms | ✓ 621ms | ✓ 984ms | ✓ 829ms | http |
| 195.158.8.123:3128 | ✓ 1660ms | 否 | ✓ 1652ms | 否 | ✓ 1730ms | http |
| 142.171.95.105:3128 | ✓ 645ms | ✓ 1403ms | ✓ 1101ms | ✓ 1071ms | ✓ 867ms | http |
| 162.240.154.26:3128 | ✓ 712ms | ✓ 1282ms | ✓ 1037ms | ✓ 1217ms | ✓ 1050ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1057ms | ✓ 1394ms | ✓ 1093ms | http |
| 103.3.246.71:3128 | ✓ 1172ms | 否 | ✓ 1086ms | ✓ 1582ms | ✓ 1102ms | http |
| 137.59.47.73:3128 | ✓ 1614ms | 否 | ✓ 1698ms | ✓ 1664ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1347ms | ✓ 1940ms | ✓ 1787ms | 否 | 否 | http |
| 8.219.97.248:80 | ✓ 1120ms | 否 | ✓ 1377ms | ✓ 1959ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1263ms | 否 | ✓ 1148ms | ✓ 1664ms | ✓ 1213ms | http |
| 212.58.132.5:8888 | ✓ 1623ms | 否 | ✓ 1489ms | ✓ 1751ms | ✓ 1247ms | http |
| 195.200.31.45:3128 | ✓ 1519ms | 否 | 否 | ✓ 1452ms | ✓ 1444ms | http |
| 46.30.46.133:3128 | ✓ 1866ms | ✓ 1231ms | ✓ 466ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1682ms | 否 | ✓ 1374ms | ✓ 1940ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1438ms | ✓ 1243ms | 否 | ✓ 1213ms | http |
| 2.78.60.10:3129 | ✓ 1367ms | 否 | ✓ 1342ms | 否 | ✓ 1827ms | http |
| 174.140.109.250:3128 | ✓ 765ms | 否 | ✓ 218ms | ✓ 1156ms | ✓ 784ms | http |
| 8.219.64.245:3128 | ✓ 943ms | 否 | ✓ 1007ms | ✓ 1333ms | ✓ 1996ms | http |
| 45.140.147.155:1082 | ✓ 1527ms | ✓ 1442ms | 否 | 否 | ✓ 1067ms | http |
| 36.141.21.200:7890 | ✓ 1097ms | ✓ 1350ms | ✓ 1185ms | ✓ 1506ms | ✓ 1150ms | http |
| 38.180.2.107:3128 | ✓ 967ms | ✓ 1971ms | ✓ 1658ms | 否 | ✓ 1903ms | http |
| 45.140.147.82:1081 | ✓ 384ms | 否 | ✓ 653ms | ✓ 1931ms | ✓ 1008ms | http |
| 222.228.171.92:8080 | ✓ 835ms | 否 | ✓ 827ms | ✓ 1283ms | ✓ 974ms | http |
| 45.140.147.82:1082 | ✓ 405ms | 否 | ✓ 600ms | 否 | ✓ 1959ms | http |
| 150.241.106.173:8080 | 否 | 否 | ✓ 557ms | ✓ 1940ms | ✓ 1526ms | http |
| 5.196.101.18:3128 | ✓ 917ms | 否 | ✓ 636ms | ✓ 1643ms | ✓ 1336ms | http |
| 173.212.246.157:3128 | ✓ 944ms | 否 | ✓ 664ms | 否 | ✓ 1567ms | http |
| 36.103.198.235:7890 | 否 | ✓ 1513ms | ✓ 1260ms | 否 | ✓ 1288ms | http |
| 223.16.170.103:3128 | ✓ 1316ms | 否 | ✓ 1296ms | 否 | ✓ 1310ms | http |
| 157.66.16.68:8686 | 否 | 否 | ✓ 1594ms | ✓ 1734ms | ✓ 1723ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1853ms | ✓ 1038ms | 否 | ✓ 1767ms | http |
| 118.31.1.154:80 | ✓ 1015ms | ✓ 1297ms | ✓ 1097ms | ✓ 1448ms | ✓ 1066ms | http |
| 114.237.77.219:1080 | 否 | ✓ 1409ms | 否 | ✓ 1745ms | ✓ 1207ms | http |
| 213.131.85.28:1976 | 否 | 否 | ✓ 1613ms | ✓ 1881ms | ✓ 1572ms | http |
| 139.28.49.90:8080 | ✓ 1132ms | ✓ 1815ms | ✓ 1784ms | 否 | 否 | http |
| 139.99.238.95:8080 | 否 | 否 | ✓ 1260ms | ✓ 1935ms | ✓ 1423ms | http |
| 209.126.10.139:3128 | ✓ 1472ms | 否 | ✓ 1052ms | ✓ 1062ms | ✓ 977ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1482ms | ✓ 1713ms | ✓ 1694ms | ✓ 1095ms | http |
| 171.227.167.109:1009 | ✓ 1913ms | 否 | ✓ 1571ms | 否 | ✓ 1170ms | http |
| 128.199.116.219:9090 | ✓ 1584ms | 否 | ✓ 1300ms | ✓ 1465ms | ✓ 1273ms | http |
| 195.26.224.49:3128 | ✓ 1103ms | 否 | ✓ 576ms | 否 | ✓ 1495ms | http |

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
