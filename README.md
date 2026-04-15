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

最后更新：2026-04-15 16:00:05 UTC（2026-04-16 00:00:05 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 657ms | 否 | ✓ 1095ms | ✓ 1040ms | ✓ 1317ms | http |
| 167.103.115.102:8800 | ✓ 922ms | 否 | ✓ 1099ms | ✓ 985ms | ✓ 953ms | http |
| 1.231.81.166:3128 | ✓ 963ms | 否 | ✓ 1262ms | ✓ 1270ms | ✓ 934ms | http |
| 113.160.132.26:8080 | ✓ 1568ms | 否 | 否 | ✓ 1284ms | ✓ 1144ms | http |
| 167.103.34.108:8800 | ✓ 1529ms | 否 | ✓ 1446ms | ✓ 1788ms | ✓ 1532ms | http |
| 157.230.178.216:8088 | ✓ 587ms | 否 | ✓ 985ms | ✓ 1517ms | ✓ 1377ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1889ms | ✓ 1171ms | ✓ 750ms | http |
| 167.103.144.127:8800 | ✓ 1544ms | 否 | ✓ 978ms | 否 | ✓ 1561ms | http |
| 159.223.225.118:8888 | ✓ 1119ms | 否 | ✓ 880ms | 否 | ✓ 1628ms | http |
| 167.103.31.122:8800 | ✓ 1835ms | 否 | ✓ 1879ms | 否 | ✓ 1877ms | http |
| 137.59.47.73:3128 | ✓ 1078ms | 否 | ✓ 1659ms | 否 | ✓ 1139ms | http |
| 45.167.125.21:999 | ✓ 1373ms | 否 | 否 | ✓ 1893ms | ✓ 1736ms | http |
| 20.27.14.220:8561 | ✓ 1501ms | ✓ 1483ms | ✓ 843ms | ✓ 997ms | ✓ 752ms | http |
| 168.110.52.228:3128 | ✓ 496ms | 否 | 否 | ✓ 757ms | ✓ 584ms | http |
| 150.107.140.238:3128 | ✓ 1831ms | 否 | ✓ 842ms | ✓ 1099ms | ✓ 901ms | http |
| 130.61.30.221:8080 | ✓ 903ms | ✓ 1829ms | ✓ 812ms | ✓ 1640ms | 否 | http |
| 103.178.171.113:8181 | ✓ 1847ms | 否 | ✓ 1612ms | ✓ 1735ms | ✓ 1386ms | http |
| 36.103.198.235:7890 | ✓ 1062ms | ✓ 1331ms | 否 | ✓ 1510ms | ✓ 1430ms | http |
| 45.12.151.226:2829 | ✓ 1973ms | ✓ 1866ms | ✓ 899ms | 否 | 否 | http |
| 85.239.59.252:7890 | ✓ 915ms | 否 | ✓ 846ms | 否 | ✓ 1505ms | http |
| 20.27.13.35:8561 | ✓ 1319ms | ✓ 1873ms | ✓ 1416ms | 否 | 否 | http |
| 20.78.26.206:8561 | ✓ 836ms | ✓ 1275ms | ✓ 1729ms | 否 | 否 | http |
| 20.210.76.104:8561 | ✓ 589ms | ✓ 938ms | ✓ 741ms | ✓ 821ms | ✓ 647ms | http |
| 20.210.76.178:8561 | ✓ 578ms | ✓ 1168ms | ✓ 619ms | ✓ 822ms | ✓ 615ms | http |
| 20.27.15.111:8561 | ✓ 1328ms | ✓ 1769ms | ✓ 1488ms | 否 | ✓ 667ms | http |
| 138.124.108.176:3128 | ✓ 779ms | 否 | ✓ 1121ms | ✓ 1858ms | ✓ 1567ms | http |
| 20.127.128.70:8080 | ✓ 1844ms | 否 | ✓ 1289ms | 否 | ✓ 1152ms | http |
| 147.161.239.240:8800 | ✓ 975ms | 否 | ✓ 1336ms | ✓ 1848ms | ✓ 1559ms | http |
| 34.101.184.164:3128 | ✓ 1554ms | 否 | ✓ 1121ms | ✓ 1689ms | ✓ 1102ms | http |
| 181.78.44.63:999 | ✓ 1721ms | ✓ 1708ms | ✓ 1337ms | ✓ 1551ms | ✓ 1524ms | http |
| 78.11.96.22:8888 | ✓ 1235ms | 否 | ✓ 1313ms | ✓ 1808ms | 否 | http |
| 35.225.22.61:80 | ✓ 960ms | 否 | ✓ 1249ms | 否 | ✓ 1328ms | http |
| 144.31.27.49:1080 | ✓ 982ms | 否 | ✓ 1434ms | 否 | ✓ 1921ms | http |
| 152.32.132.190:7890 | ✓ 1518ms | ✓ 1899ms | ✓ 1555ms | ✓ 1030ms | ✓ 855ms | http |
| 185.132.178.178:1080 | 否 | 否 | ✓ 1223ms | ✓ 1368ms | ✓ 1440ms | http |
| 177.234.217.88:999 | ✓ 1443ms | 否 | ✓ 1223ms | ✓ 1911ms | ✓ 1563ms | http |
| 20.210.39.153:8561 | ✓ 444ms | ✓ 961ms | ✓ 443ms | ✓ 793ms | ✓ 652ms | http |
| 20.78.118.91:8561 | ✓ 465ms | ✓ 1039ms | ✓ 438ms | ✓ 751ms | ✓ 629ms | http |
| 8.219.97.248:80 | ✓ 1133ms | 否 | ✓ 903ms | ✓ 1131ms | 否 | http |
| 36.141.21.200:7890 | ✓ 1141ms | ✓ 1082ms | ✓ 1042ms | ✓ 1195ms | ✓ 962ms | http |
| 20.27.11.248:8561 | ✓ 971ms | 否 | ✓ 990ms | ✓ 1170ms | ✓ 778ms | http |
| 157.0.142.246:10061 | ✓ 1914ms | ✓ 1311ms | ✓ 1080ms | ✓ 1347ms | 否 | http |
| 94.131.118.129:1081 | ✓ 1280ms | 否 | ✓ 1839ms | 否 | ✓ 1623ms | http |
| 120.92.108.86:7890 | ✓ 1451ms | 否 | ✓ 1231ms | 否 | ✓ 1818ms | http |
| 164.163.40.1:10000 | ✓ 1391ms | 否 | ✓ 1569ms | 否 | ✓ 1998ms | http |
| 140.238.242.189:8100 | ✓ 1743ms | 否 | 否 | ✓ 1939ms | ✓ 1454ms | http |
| 107.172.102.234:40621 | 否 | ✓ 1233ms | ✓ 681ms | ✓ 1277ms | ✓ 585ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1311ms | 否 | ✓ 1412ms | ✓ 1567ms | http |
| 45.140.147.82:1081 | ✓ 885ms | ✓ 1436ms | ✓ 1059ms | 否 | ✓ 1354ms | http |
| 116.80.64.157:7777 | ✓ 1526ms | 否 | ✓ 1460ms | ✓ 1792ms | 否 | http |
| 162.240.154.26:3128 | ✓ 764ms | ✓ 1879ms | ✓ 357ms | ✓ 814ms | ✓ 621ms | http |
| 194.87.85.207:1080 | ✓ 1179ms | 否 | ✓ 1209ms | 否 | ✓ 1666ms | http |
| 147.45.186.28:3128 | ✓ 1383ms | 否 | ✓ 1601ms | 否 | ✓ 1596ms | http |
| 45.140.147.155:1082 | ✓ 637ms | ✓ 1398ms | ✓ 1322ms | ✓ 1727ms | 否 | http |
| 150.241.116.228:3128 | ✓ 1373ms | 否 | ✓ 1342ms | 否 | ✓ 1902ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1612ms | ✓ 1626ms | ✓ 1065ms | http |
| 104.248.211.46:7890 | ✓ 629ms | ✓ 1543ms | 否 | ✓ 1647ms | ✓ 1986ms | http |
| 27.71.24.102:3128 | 否 | 否 | ✓ 1178ms | ✓ 1078ms | ✓ 823ms | http |
| 77.110.113.24:40000 | ✓ 1157ms | 否 | ✓ 1924ms | 否 | ✓ 1497ms | http |
| 195.26.224.49:3128 | ✓ 857ms | 否 | ✓ 930ms | ✓ 1887ms | ✓ 1588ms | http |
| 120.92.212.16:7890 | ✓ 1098ms | ✓ 1245ms | ✓ 1613ms | ✓ 1810ms | ✓ 952ms | http |
| 121.230.8.111:1080 | ✓ 1562ms | 否 | ✓ 1592ms | 否 | ✓ 1866ms | http |
| 185.114.73.2:1080 | ✓ 1448ms | 否 | ✓ 1566ms | ✓ 1785ms | ✓ 1954ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1936ms | ✓ 1627ms | 否 | ✓ 1859ms | http |
| 45.140.147.82:1082 | ✓ 1538ms | 否 | ✓ 704ms | 否 | ✓ 1004ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1229ms | ✓ 1256ms | ✓ 1078ms | http |
| 49.156.44.10:8080 | 否 | 否 | ✓ 1197ms | ✓ 1269ms | ✓ 1276ms | http |
| 116.80.64.41:7777 | ✓ 1505ms | 否 | ✓ 1993ms | 否 | ✓ 1617ms | http |
| 212.58.132.5:8888 | ✓ 1097ms | 否 | ✓ 1669ms | ✓ 1455ms | 否 | http |
| 196.217.1.63:1221 | ✓ 1589ms | 否 | ✓ 1915ms | 否 | ✓ 1936ms | http |
| 218.108.131.186:17890 | ✓ 1010ms | 否 | ✓ 807ms | ✓ 1061ms | ✓ 1097ms | http |
| 202.141.161.53:10808 | ✓ 1019ms | ✓ 1936ms | 否 | ✓ 1722ms | 否 | http |
| 120.92.212.16:8890 | ✓ 990ms | ✓ 1919ms | ✓ 1695ms | ✓ 1193ms | ✓ 1792ms | http |
| 59.46.216.131:30001 | ✓ 1591ms | ✓ 1504ms | 否 | 否 | ✓ 1949ms | http |
| 185.191.236.162:3128 | ✓ 1123ms | 否 | ✓ 1690ms | ✓ 1952ms | ✓ 1911ms | http |
| 94.131.118.39:1081 | ✓ 643ms | 否 | ✓ 1724ms | 否 | ✓ 1622ms | http |
| 94.131.118.39:1082 | ✓ 661ms | 否 | ✓ 1703ms | 否 | ✓ 1622ms | http |
| 62.113.119.14:8080 | ✓ 810ms | ✓ 1753ms | ✓ 1693ms | ✓ 1894ms | ✓ 1422ms | http |
| 121.148.239.82:3052 | ✓ 764ms | ✓ 1352ms | ✓ 1589ms | ✓ 953ms | ✓ 1032ms | http |
| 181.188.203.104:999 | ✓ 1728ms | 否 | 否 | ✓ 1922ms | ✓ 1857ms | http |
| 201.182.150.184:999 | ✓ 1582ms | 否 | 否 | ✓ 1946ms | ✓ 1895ms | http |
| 61.52.131.172:8443 | ✓ 881ms | ✓ 1173ms | ✓ 917ms | ✓ 1200ms | ✓ 981ms | http |
| 124.83.112.246:8082 | ✓ 1292ms | 否 | ✓ 1092ms | ✓ 1736ms | ✓ 1433ms | http |
| 103.113.70.189:1081 | ✓ 598ms | ✓ 1887ms | ✓ 1285ms | ✓ 1191ms | ✓ 906ms | http |

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
