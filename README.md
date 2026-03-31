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

最后更新：2026-03-31 10:03:18 UTC（2026-03-31 18:03:18 UTC+8）

**代理总数：91**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.239.240:8800 | ✓ 733ms | ✓ 1432ms | ✓ 885ms | ✓ 1582ms | ✓ 1303ms | http |
| 39.185.46.193:5911 | ✓ 1026ms | ✓ 1013ms | ✓ 915ms | ✓ 1470ms | ✓ 974ms | http |
| 95.213.217.168:52004 | ✓ 764ms | ✓ 1606ms | ✓ 1737ms | 否 | ✓ 1585ms | http |
| 167.103.115.102:8800 | ✓ 1176ms | 否 | ✓ 1128ms | ✓ 1251ms | ✓ 1322ms | http |
| 1.231.81.166:3128 | ✓ 1822ms | ✓ 1419ms | ✓ 1789ms | ✓ 1216ms | ✓ 1008ms | http |
| 147.161.210.140:8800 | ✓ 1790ms | 否 | ✓ 950ms | ✓ 1822ms | ✓ 1063ms | http |
| 45.136.198.40:3128 | ✓ 758ms | 否 | ✓ 1683ms | ✓ 1792ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1875ms | ✓ 1857ms | ✓ 1396ms | ✓ 1455ms | ✓ 1373ms | http |
| 167.103.34.108:8800 | ✓ 1591ms | 否 | ✓ 1734ms | ✓ 1863ms | ✓ 1607ms | http |
| 42.96.16.158:1311 | ✓ 1871ms | 否 | 否 | ✓ 1454ms | ✓ 1141ms | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1200ms | ✓ 1806ms | ✓ 1544ms | http |
| 35.225.22.61:80 | ✓ 532ms | ✓ 1485ms | 否 | ✓ 1071ms | 否 | http |
| 91.238.123.111:8000 | ✓ 1635ms | 否 | ✓ 844ms | ✓ 1535ms | ✓ 990ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1213ms | ✓ 1830ms | ✓ 1352ms | http |
| 167.103.144.127:8800 | ✓ 1720ms | 否 | ✓ 1550ms | 否 | ✓ 1576ms | http |
| 91.238.123.230:8000 | 否 | 否 | ✓ 1650ms | ✓ 1509ms | ✓ 1207ms | http |
| 185.191.236.162:3128 | 否 | 否 | ✓ 941ms | ✓ 1888ms | ✓ 1187ms | http |
| 43.99.54.236:5555 | ✓ 829ms | ✓ 1218ms | ✓ 842ms | ✓ 1040ms | ✓ 830ms | http |
| 62.113.119.14:8080 | ✓ 1377ms | 否 | ✓ 1789ms | 否 | ✓ 1727ms | http |
| 167.103.31.122:8800 | ✓ 1453ms | 否 | ✓ 1313ms | 否 | ✓ 1693ms | http |
| 208.87.243.199:7878 | ✓ 384ms | ✓ 888ms | ✓ 909ms | ✓ 1131ms | ✓ 1017ms | http |
| 209.126.84.232:8888 | ✓ 1126ms | ✓ 1610ms | ✓ 301ms | ✓ 1830ms | ✓ 941ms | http |
| 38.34.179.79:8451 | ✓ 1041ms | ✓ 935ms | ✓ 1313ms | ✓ 989ms | ✓ 1309ms | http |
| 86.53.183.16:1080 | ✓ 727ms | ✓ 1469ms | 否 | 否 | ✓ 1201ms | http |
| 133.242.138.34:8100 | ✓ 1618ms | ✓ 1296ms | ✓ 998ms | ✓ 1570ms | ✓ 1177ms | http |
| 103.84.95.54:7890 | ✓ 952ms | 否 | ✓ 889ms | ✓ 1774ms | 否 | http |
| 150.241.71.15:1080 | 否 | ✓ 1755ms | ✓ 479ms | 否 | ✓ 1984ms | http |
| 160.238.65.9:3128 | ✓ 495ms | ✓ 1432ms | ✓ 1557ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 1044ms | ✓ 1349ms | ✓ 1140ms | ✓ 1422ms | ✓ 1083ms | http |
| 5.104.87.17:8051 | ✓ 1507ms | 否 | ✓ 1582ms | ✓ 1993ms | 否 | http |
| 160.238.65.2:3128 | ✓ 488ms | ✓ 1309ms | ✓ 1684ms | 否 | ✓ 1535ms | http |
| 120.92.212.16:8890 | ✓ 1240ms | 否 | 否 | ✓ 1461ms | ✓ 1582ms | http |
| 45.12.151.226:2829 | ✓ 874ms | ✓ 1654ms | ✓ 1503ms | 否 | 否 | http |
| 38.34.179.83:8448 | ✓ 1566ms | 否 | ✓ 1673ms | 否 | ✓ 1827ms | http |
| 160.238.65.7:3128 | ✓ 487ms | ✓ 1685ms | ✓ 1309ms | 否 | ✓ 1538ms | http |
| 177.234.217.88:999 | ✓ 1479ms | ✓ 1980ms | ✓ 1786ms | ✓ 1961ms | ✓ 1695ms | http |
| 31.192.106.135:8010 | ✓ 953ms | 否 | 否 | ✓ 1599ms | ✓ 1743ms | http |
| 217.217.249.160:8080 | ✓ 1694ms | 否 | ✓ 1169ms | 否 | ✓ 1217ms | http |
| 222.228.171.92:8080 | ✓ 1791ms | 否 | ✓ 838ms | ✓ 1260ms | ✓ 1068ms | http |
| 20.78.213.56:80 | ✓ 826ms | ✓ 1579ms | 否 | ✓ 1247ms | ✓ 1001ms | http |
| 103.82.23.118:5226 | ✓ 1683ms | 否 | ✓ 1532ms | 否 | ✓ 1729ms | http |
| 181.78.44.63:999 | 否 | 否 | ✓ 1369ms | ✓ 1901ms | ✓ 1575ms | http |
| 120.92.212.16:7890 | ✓ 1344ms | 否 | ✓ 1285ms | ✓ 1997ms | ✓ 1618ms | http |
| 121.126.185.63:25152 | ✓ 1664ms | 否 | 否 | ✓ 1953ms | ✓ 1625ms | http |
| 210.223.44.230:3128 | ✓ 1353ms | ✓ 1452ms | ✓ 1168ms | ✓ 1178ms | ✓ 984ms | http |
| 158.160.215.167:8126 | 否 | ✓ 1981ms | ✓ 915ms | 否 | ✓ 1706ms | http |
| 158.160.215.167:8127 | ✓ 701ms | ✓ 1766ms | ✓ 1053ms | ✓ 1895ms | ✓ 1916ms | http |
| 34.101.184.164:3128 | ✓ 1656ms | 否 | ✓ 991ms | ✓ 1739ms | ✓ 1695ms | http |
| 103.133.254.4:3128 | ✓ 1936ms | 否 | ✓ 1710ms | 否 | ✓ 1706ms | http |
| 125.76.214.178:8091 | ✓ 1318ms | ✓ 1570ms | ✓ 1463ms | ✓ 1547ms | ✓ 1221ms | http |
| 201.150.116.3:999 | ✓ 1952ms | ✓ 1507ms | ✓ 1286ms | 否 | 否 | http |
| 23.253.80.88:80 | ✓ 325ms | 否 | ✓ 1892ms | ✓ 1702ms | 否 | http |
| 106.10.55.212:1121 | ✓ 1157ms | ✓ 1433ms | ✓ 1240ms | 否 | 否 | http |
| 103.97.88.47:3128 | ✓ 1209ms | ✓ 1860ms | ✓ 1930ms | 否 | ✓ 1612ms | http |
| 38.51.232.90:1986 | ✓ 1009ms | ✓ 1473ms | ✓ 1355ms | 否 | 否 | http |
| 103.173.139.221:8080 | ✓ 1619ms | 否 | ✓ 1533ms | ✓ 1973ms | ✓ 1801ms | http |
| 150.107.140.238:3128 | ✓ 1882ms | 否 | ✓ 1380ms | ✓ 1533ms | ✓ 1135ms | http |
| 158.160.215.167:8124 | ✓ 708ms | 否 | ✓ 1167ms | 否 | ✓ 1771ms | http |
| 103.82.93.100:3128 | ✓ 1863ms | 否 | ✓ 1273ms | ✓ 1589ms | 否 | http |
| 65.21.201.149:8080 | ✓ 1775ms | ✓ 1856ms | ✓ 1129ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 396ms | ✓ 1842ms | ✓ 875ms | ✓ 1293ms | ✓ 727ms | http |
| 59.46.216.131:30001 | ✓ 1170ms | ✓ 1579ms | ✓ 1794ms | ✓ 1560ms | 否 | http |
| 167.71.196.28:8080 | ✓ 984ms | 否 | ✓ 1808ms | 否 | ✓ 1109ms | http |
| 46.39.105.157:8080 | ✓ 1069ms | 否 | ✓ 1323ms | ✓ 1818ms | ✓ 1483ms | http |
| 217.217.249.160:80 | ✓ 983ms | 否 | ✓ 1044ms | ✓ 1897ms | ✓ 1720ms | http |
| 65.108.203.37:18080 | ✓ 1184ms | 否 | ✓ 1127ms | 否 | ✓ 1230ms | http |
| 160.238.65.6:3128 | ✓ 1838ms | ✓ 1431ms | ✓ 408ms | ✓ 1250ms | ✓ 941ms | http |
| 160.238.65.3:3128 | ✓ 1836ms | 否 | ✓ 400ms | ✓ 1246ms | ✓ 946ms | http |
| 200.174.198.32:8888 | ✓ 1432ms | 否 | ✓ 690ms | 否 | ✓ 1898ms | http |
| 103.39.51.190:8080 | ✓ 1921ms | 否 | ✓ 1818ms | ✓ 1522ms | ✓ 1559ms | http |
| 197.164.101.11:1981 | 否 | 否 | ✓ 1470ms | ✓ 1952ms | ✓ 1477ms | http |
| 103.52.115.171:3128 | 否 | 否 | ✓ 1002ms | ✓ 1521ms | ✓ 1221ms | http |
| 34.96.238.40:8080 | ✓ 1396ms | 否 | ✓ 1214ms | ✓ 1826ms | 否 | http |
| 38.145.218.13:8451 | ✓ 913ms | ✓ 1296ms | ✓ 1470ms | ✓ 1204ms | ✓ 1965ms | http |
| 103.82.23.118:5221 | ✓ 1630ms | ✓ 1748ms | ✓ 1524ms | 否 | ✓ 1837ms | http |
| 38.194.231.66:999 | 否 | ✓ 1279ms | ✓ 1224ms | 否 | ✓ 1264ms | http |
| 197.164.101.13:1976 | 否 | 否 | ✓ 1753ms | ✓ 1936ms | ✓ 1506ms | http |
| 121.230.9.185:1080 | 否 | ✓ 1968ms | ✓ 1486ms | 否 | ✓ 1672ms | http |
| 38.145.218.134:8445 | ✓ 621ms | ✓ 1790ms | ✓ 1275ms | 否 | 否 | http |
| 121.230.8.235:1080 | ✓ 1359ms | ✓ 1980ms | ✓ 1376ms | 否 | 否 | http |
| 198.59.68.130:3128 | ✓ 951ms | ✓ 1627ms | ✓ 1736ms | 否 | ✓ 1211ms | http |
| 157.66.2.100:1111 | ✓ 1510ms | 否 | ✓ 1942ms | ✓ 1684ms | ✓ 1666ms | http |
| 194.147.149.234:3128 | ✓ 488ms | 否 | ✓ 1251ms | ✓ 1605ms | 否 | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1981ms | ✓ 1988ms | ✓ 1995ms | http |
| 45.136.130.246:8446 | ✓ 906ms | ✓ 1281ms | ✓ 727ms | ✓ 968ms | ✓ 994ms | http |
| 5.104.87.17:8050 | ✓ 1187ms | ✓ 1925ms | ✓ 1733ms | ✓ 1229ms | 否 | http |
| 47.95.231.180:8084 | ✓ 1051ms | ✓ 1403ms | ✓ 1156ms | ✓ 1576ms | ✓ 1468ms | http |
| 138.197.68.35:4857 | ✓ 386ms | 否 | ✓ 63ms | ✓ 1085ms | 否 | http |
| 103.118.102.98:80 | ✓ 1951ms | 否 | 否 | ✓ 1633ms | ✓ 1870ms | http |
| 157.230.220.25:4857 | ✓ 128ms | ✓ 1537ms | 否 | 否 | ✓ 764ms | http |
| 118.31.1.154:80 | ✓ 1085ms | ✓ 1329ms | ✓ 1077ms | ✓ 1332ms | ✓ 1055ms | http |

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
