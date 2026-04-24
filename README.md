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

最后更新：2026-04-24 23:37:14 UTC（2026-04-25 07:37:14 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.211.166.184:8081 | ✓ 693ms | ✓ 1283ms | ✓ 905ms | ✓ 1085ms | ✓ 863ms | http |
| 46.101.95.183:8888 | ✓ 774ms | 否 | ✓ 543ms | ✓ 1995ms | ✓ 1585ms | http |
| 2.27.54.161:1080 | ✓ 952ms | 否 | ✓ 1820ms | 否 | ✓ 1390ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1856ms | ✓ 1824ms | ✓ 1816ms | ✓ 1536ms | http |
| 35.225.22.61:80 | ✓ 408ms | ✓ 1345ms | ✓ 1054ms | 否 | ✓ 1116ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1569ms | ✓ 1508ms | ✓ 1588ms | ✓ 1167ms | http |
| 47.85.51.197:1080 | ✓ 885ms | 否 | ✓ 1504ms | ✓ 1073ms | 否 | http |
| 217.76.245.80:999 | ✓ 822ms | ✓ 1550ms | ✓ 1166ms | ✓ 1548ms | ✓ 1543ms | http |
| 168.144.75.9:3128 | 否 | 否 | ✓ 1594ms | ✓ 1887ms | ✓ 1462ms | http |
| 46.249.98.211:3128 | ✓ 697ms | ✓ 1501ms | ✓ 1982ms | 否 | 否 | http |
| 106.44.155.90:2222 | ✓ 1178ms | ✓ 1596ms | ✓ 1241ms | ✓ 1645ms | ✓ 1377ms | http |
| 45.140.147.82:1082 | ✓ 883ms | ✓ 1182ms | ✓ 421ms | ✓ 1379ms | ✓ 1981ms | http |
| 3.19.213.118:40000 | ✓ 1062ms | 否 | ✓ 1453ms | ✓ 1590ms | ✓ 1123ms | http |
| 15.237.108.20:994 | ✓ 1060ms | 否 | ✓ 787ms | ✓ 1674ms | ✓ 1221ms | http |
| 18.188.53.175:17723 | ✓ 1089ms | 否 | ✓ 878ms | ✓ 1693ms | 否 | http |
| 18.222.132.180:35576 | ✓ 1075ms | 否 | ✓ 1716ms | ✓ 1743ms | ✓ 1399ms | http |
| 34.246.223.187:4192 | 否 | 否 | ✓ 549ms | ✓ 1971ms | ✓ 1500ms | http |
| 54.229.201.146:48867 | ✓ 1052ms | 否 | ✓ 1371ms | ✓ 1930ms | ✓ 1651ms | http |
| 91.217.81.131:1080 | ✓ 1060ms | ✓ 1751ms | ✓ 1662ms | 否 | ✓ 1521ms | http |
| 52.16.215.4:42899 | ✓ 1754ms | 否 | ✓ 1528ms | 否 | ✓ 1508ms | http |
| 51.17.5.6:1279 | ✓ 1748ms | 否 | ✓ 973ms | 否 | ✓ 1703ms | http |
| 51.17.154.141:69 | ✓ 1140ms | 否 | ✓ 1731ms | 否 | ✓ 1749ms | http |
| 130.61.174.200:1080 | ✓ 1029ms | ✓ 1539ms | ✓ 1880ms | ✓ 1868ms | 否 | http |
| 161.35.181.96:999 | ✓ 737ms | ✓ 1188ms | ✓ 829ms | ✓ 883ms | ✓ 936ms | http |
| 34.71.229.255:3128 | ✓ 930ms | ✓ 1869ms | ✓ 1725ms | 否 | 否 | http |
| 45.140.147.82:1081 | ✓ 438ms | ✓ 1296ms | ✓ 1000ms | 否 | 否 | http |
| 160.238.65.2:3128 | 否 | ✓ 1829ms | ✓ 617ms | ✓ 1782ms | 否 | http |
| 59.46.216.131:30001 | 否 | ✓ 1616ms | ✓ 1324ms | ✓ 1490ms | 否 | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 978ms | ✓ 1256ms | ✓ 976ms | http |
| 160.250.4.245:1 | 否 | 否 | ✓ 1331ms | ✓ 1771ms | ✓ 1186ms | http |
| 177.93.132.244:3128 | ✓ 647ms | 否 | ✓ 1281ms | 否 | ✓ 1698ms | http |
| 82.148.18.242:443 | ✓ 636ms | 否 | ✓ 1961ms | ✓ 1585ms | 否 | http |
| 91.99.15.45:2095 | ✓ 1221ms | 否 | ✓ 1694ms | ✓ 1445ms | 否 | http |
| 178.156.224.42:3128 | ✓ 1091ms | 否 | ✓ 1600ms | 否 | ✓ 1496ms | http |
| 160.238.65.3:3128 | ✓ 1537ms | 否 | 否 | ✓ 1589ms | ✓ 1621ms | http |
| 101.32.244.83:8080 | ✓ 1150ms | ✓ 1641ms | ✓ 1234ms | ✓ 1674ms | ✓ 1496ms | http |
| 121.43.196.213:8222 | ✓ 1144ms | ✓ 1260ms | ✓ 1172ms | ✓ 1343ms | ✓ 1057ms | http |
| 121.43.196.210:8222 | ✓ 1184ms | ✓ 1382ms | ✓ 1017ms | ✓ 1313ms | ✓ 1059ms | http |
| 114.55.226.123:10086 | ✓ 1257ms | ✓ 1679ms | ✓ 1173ms | ✓ 1617ms | ✓ 1242ms | http |
| 108.130.79.116:8999 | ✓ 873ms | 否 | ✓ 1167ms | ✓ 1801ms | ✓ 1499ms | http |
| 15.157.63.22:15433 | ✓ 953ms | 否 | ✓ 1183ms | ✓ 1918ms | ✓ 1811ms | http |
| 13.38.59.232:45664 | ✓ 647ms | 否 | ✓ 1157ms | ✓ 1940ms | ✓ 1541ms | http |
| 52.47.115.41:8698 | 否 | 否 | ✓ 605ms | ✓ 1619ms | ✓ 1622ms | http |
| 13.41.196.179:2130 | ✓ 757ms | 否 | ✓ 985ms | 否 | ✓ 1154ms | http |
| 16.62.123.236:48789 | ✓ 1153ms | 否 | ✓ 667ms | ✓ 1736ms | ✓ 1211ms | http |
| 15.160.116.45:13815 | ✓ 905ms | 否 | ✓ 996ms | ✓ 1820ms | ✓ 1336ms | http |
| 42.101.8.101:8888 | ✓ 1336ms | 否 | ✓ 1509ms | ✓ 1624ms | ✓ 1388ms | http |
| 116.80.91.222:3172 | ✓ 1631ms | ✓ 1220ms | ✓ 781ms | ✓ 997ms | ✓ 825ms | http |
| 120.27.224.64:3128 | ✓ 1034ms | ✓ 1336ms | ✓ 1069ms | ✓ 1312ms | ✓ 1093ms | http |
| 103.157.200.126:3128 | ✓ 1200ms | 否 | ✓ 1162ms | ✓ 1574ms | ✓ 1538ms | http |
| 91.193.240.157:9877 | ✓ 1353ms | 否 | ✓ 1284ms | 否 | ✓ 1954ms | http |
| 43.132.188.134:443 | ✓ 1709ms | ✓ 1777ms | 否 | 否 | ✓ 1038ms | http |
| 117.236.124.166:3128 | ✓ 1900ms | 否 | ✓ 1395ms | ✓ 1731ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1923ms | 否 | 否 | ✓ 1821ms | ✓ 1724ms | http |
| 120.92.212.16:8890 | ✓ 1932ms | ✓ 1381ms | ✓ 1823ms | 否 | ✓ 1749ms | http |
| 218.108.131.186:17890 | ✓ 995ms | ✓ 1259ms | ✓ 1055ms | ✓ 1348ms | ✓ 1027ms | http |
| 60.249.94.208:3128 | ✓ 1164ms | 否 | ✓ 995ms | ✓ 1196ms | ✓ 970ms | http |
| 183.232.248.73:7890 | ✓ 1023ms | ✓ 1326ms | ✓ 1122ms | ✓ 1272ms | ✓ 1004ms | http |
| 160.238.65.7:3128 | ✓ 1563ms | ✓ 1685ms | ✓ 1745ms | 否 | 否 | http |
| 160.238.65.5:3128 | 否 | ✓ 1488ms | 否 | ✓ 1896ms | ✓ 1045ms | http |
| 84.47.150.125:1080 | ✓ 961ms | 否 | 否 | ✓ 1999ms | ✓ 1629ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1581ms | ✓ 1736ms | ✓ 1345ms | 否 | http |
| 108.130.79.116:51858 | 否 | 否 | ✓ 471ms | ✓ 1790ms | ✓ 1053ms | http |
| 15.160.132.166:8800 | ✓ 824ms | 否 | ✓ 1237ms | 否 | ✓ 1580ms | http |
| 35.178.250.178:8888 | ✓ 819ms | 否 | 否 | ✓ 1818ms | ✓ 1287ms | http |
| 45.129.141.143:3128 | ✓ 926ms | ✓ 1798ms | ✓ 1600ms | ✓ 1786ms | ✓ 1675ms | http |
| 45.140.147.155:1082 | ✓ 803ms | 否 | ✓ 1789ms | 否 | ✓ 1948ms | http |
| 47.101.159.19:8899 | ✓ 1046ms | ✓ 1280ms | ✓ 1063ms | ✓ 1315ms | ✓ 1060ms | http |
| 143.198.211.194:8080 | ✓ 1532ms | 否 | ✓ 1006ms | ✓ 1420ms | ✓ 1033ms | http |
| 47.110.42.192:9003 | ✓ 1722ms | ✓ 1862ms | ✓ 1558ms | ✓ 1750ms | ✓ 1853ms | http |
| 217.77.102.18:3128 | 否 | 否 | ✓ 1021ms | ✓ 1962ms | ✓ 1799ms | http |
| 192.99.44.178:3128 | 否 | ✓ 1747ms | ✓ 1918ms | ✓ 1906ms | 否 | http |
| 8.219.195.129:1080 | 否 | ✓ 1892ms | ✓ 881ms | ✓ 1284ms | ✓ 995ms | http |
| 149.248.76.55:10000 | 否 | 否 | ✓ 1018ms | ✓ 1741ms | ✓ 1597ms | http |
| 62.113.119.14:8080 | ✓ 820ms | ✓ 1406ms | ✓ 1329ms | ✓ 1901ms | ✓ 1184ms | http |
| 160.238.65.9:3128 | ✓ 1246ms | ✓ 1316ms | ✓ 1550ms | 否 | 否 | http |
| 160.238.65.6:3128 | ✓ 1275ms | ✓ 1309ms | ✓ 442ms | ✓ 1632ms | ✓ 1013ms | http |
| 160.238.65.4:3128 | ✓ 1276ms | ✓ 1223ms | ✓ 439ms | ✓ 1497ms | ✓ 1329ms | http |
| 152.70.91.193:40000 | 否 | 否 | ✓ 1662ms | ✓ 1927ms | ✓ 1685ms | http |
| 107.173.52.221:3128 | 否 | 否 | ✓ 1754ms | ✓ 1797ms | ✓ 1531ms | http |
| 61.52.131.172:8443 | ✓ 1631ms | ✓ 1414ms | ✓ 1147ms | ✓ 1960ms | 否 | http |
| 44.201.32.14:24330 | ✓ 1164ms | 否 | ✓ 1645ms | ✓ 1970ms | ✓ 1271ms | http |
| 103.39.51.207:8080 | ✓ 1659ms | 否 | ✓ 1877ms | 否 | ✓ 1922ms | http |

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
