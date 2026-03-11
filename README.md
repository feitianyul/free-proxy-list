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

最后更新：2026-03-11 23:26:56 UTC（2026-03-12 07:26:56 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 749ms | ✓ 976ms | ✓ 195ms | ✓ 814ms | ✓ 616ms | http |
| 45.136.130.175:8443 | ✓ 750ms | ✓ 746ms | ✓ 852ms | ✓ 808ms | ✓ 620ms | http |
| 45.136.131.47:8443 | ✓ 749ms | ✓ 942ms | ✓ 853ms | ✓ 826ms | ✓ 619ms | http |
| 1.231.81.166:3128 | ✓ 1314ms | ✓ 1000ms | ✓ 997ms | ✓ 1072ms | ✓ 766ms | http |
| 160.238.65.7:3128 | ✓ 1231ms | 否 | 否 | ✓ 1330ms | ✓ 1005ms | http |
| 202.155.12.161:443 | ✓ 1915ms | 否 | 否 | ✓ 1862ms | ✓ 1620ms | http |
| 171.251.172.78:5109 | ✓ 1631ms | 否 | ✓ 1619ms | 否 | ✓ 1483ms | http |
| 205.209.118.30:3138 | ✓ 398ms | 否 | ✓ 1032ms | ✓ 1188ms | ✓ 915ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1336ms | ✓ 1001ms | ✓ 736ms | http |
| 115.231.181.40:8128 | ✓ 922ms | 否 | ✓ 1036ms | ✓ 1185ms | ✓ 1023ms | http |
| 34.101.184.164:3128 | ✓ 1713ms | 否 | ✓ 1757ms | ✓ 1500ms | ✓ 1170ms | http |
| 152.42.213.210:8080 | ✓ 1460ms | 否 | ✓ 1784ms | ✓ 1154ms | ✓ 902ms | http |
| 107.173.52.58:7890 | ✓ 1451ms | ✓ 1305ms | ✓ 815ms | 否 | ✓ 877ms | http |
| 113.160.132.26:8080 | ✓ 1720ms | ✓ 1400ms | ✓ 1158ms | ✓ 1244ms | ✓ 1020ms | http |
| 8.219.97.248:80 | ✓ 1625ms | 否 | ✓ 1147ms | 否 | ✓ 1987ms | http |
| 86.53.183.16:1080 | ✓ 936ms | 否 | ✓ 1569ms | 否 | ✓ 1917ms | http |
| 45.136.130.188:8443 | 否 | 否 | ✓ 1643ms | ✓ 1723ms | ✓ 1110ms | http |
| 81.70.169.194:80 | 否 | ✓ 1967ms | ✓ 1646ms | 否 | ✓ 1396ms | http |
| 160.238.65.4:3128 | ✓ 1548ms | 否 | ✓ 1619ms | ✓ 1966ms | ✓ 1520ms | http |
| 160.238.65.5:3128 | 否 | ✓ 1610ms | ✓ 1551ms | ✓ 1957ms | 否 | http |
| 160.238.65.6:3128 | ✓ 1320ms | 否 | ✓ 1845ms | ✓ 1984ms | ✓ 1534ms | http |
| 190.9.109.198:999 | ✓ 817ms | ✓ 1316ms | ✓ 1226ms | ✓ 1593ms | ✓ 1114ms | http |
| 160.238.65.9:3128 | ✓ 1702ms | 否 | ✓ 1462ms | 否 | ✓ 1505ms | http |
| 46.183.25.8:443 | ✓ 980ms | 否 | ✓ 583ms | ✓ 1201ms | 否 | http |
| 45.136.130.191:8443 | ✓ 395ms | ✓ 1807ms | ✓ 434ms | ✓ 1040ms | ✓ 830ms | http |
| 45.136.130.223:8443 | ✓ 387ms | ✓ 1043ms | ✓ 976ms | ✓ 1339ms | ✓ 918ms | http |
| 80.242.56.115:3128 | ✓ 677ms | ✓ 1818ms | ✓ 645ms | ✓ 1578ms | ✓ 1634ms | http |
| 107.155.65.87:13428 | 否 | 否 | ✓ 1271ms | ✓ 1512ms | ✓ 1088ms | http |
| 170.78.208.245:999 | ✓ 513ms | 否 | 否 | ✓ 1517ms | ✓ 1576ms | http |
| 162.240.154.26:3128 | ✓ 820ms | ✓ 1447ms | ✓ 850ms | 否 | ✓ 1394ms | http |
| 1.225.116.115:1080 | ✓ 1485ms | ✓ 1677ms | 否 | ✓ 1859ms | ✓ 1519ms | http |
| 101.43.255.96:80 | ✓ 1104ms | ✓ 1477ms | ✓ 1088ms | ✓ 1279ms | ✓ 1897ms | http |
| 152.42.213.210:443 | ✓ 1387ms | 否 | 否 | ✓ 1295ms | ✓ 1257ms | http |
| 35.225.22.61:80 | ✓ 512ms | ✓ 1664ms | ✓ 885ms | ✓ 1286ms | 否 | http |
| 116.80.49.169:3172 | ✓ 1967ms | 否 | 否 | ✓ 1917ms | ✓ 1756ms | http |
| 111.48.191.1:7890 | ✓ 883ms | ✓ 1006ms | ✓ 929ms | ✓ 1071ms | ✓ 814ms | http |
| 103.35.188.243:3128 | ✓ 251ms | ✓ 1090ms | 否 | ✓ 1223ms | 否 | http |
| 195.158.8.123:3128 | ✓ 1873ms | 否 | ✓ 1766ms | 否 | ✓ 1770ms | http |
| 194.213.18.200:443 | ✓ 1035ms | ✓ 1229ms | ✓ 1096ms | ✓ 1545ms | 否 | http |
| 45.174.77.1:999 | ✓ 920ms | ✓ 1308ms | ✓ 1341ms | ✓ 1209ms | ✓ 1146ms | http |
| 84.247.149.172:3128 | ✓ 864ms | 否 | ✓ 1735ms | ✓ 1185ms | ✓ 953ms | http |
| 150.107.140.238:3128 | ✓ 1588ms | 否 | ✓ 1022ms | ✓ 1237ms | ✓ 946ms | http |
| 139.159.99.242:8080 | ✓ 970ms | ✓ 1068ms | ✓ 911ms | ✓ 1171ms | ✓ 896ms | http |
| 95.3.9.78:8080 | ✓ 791ms | ✓ 1828ms | ✓ 1825ms | ✓ 1759ms | ✓ 1358ms | http |
| 121.204.158.249:3128 | ✓ 941ms | ✓ 1230ms | ✓ 1070ms | ✓ 1267ms | ✓ 1044ms | http |
| 116.80.82.218:3172 | ✓ 1755ms | 否 | 否 | ✓ 1894ms | ✓ 1720ms | http |
| 95.3.9.78:3128 | ✓ 795ms | 否 | ✓ 749ms | ✓ 1707ms | ✓ 1294ms | http |
| 192.73.243.98:3128 | ✓ 318ms | ✓ 1096ms | ✓ 953ms | ✓ 1719ms | ✓ 1095ms | http |
| 160.238.65.3:3128 | 否 | 否 | ✓ 1086ms | ✓ 1896ms | ✓ 1023ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1798ms | ✓ 1880ms | ✓ 649ms | http |
| 160.238.65.2:3128 | 否 | ✓ 1758ms | ✓ 1324ms | ✓ 1882ms | ✓ 1022ms | http |
| 160.238.65.8:3128 | 否 | 否 | ✓ 1088ms | ✓ 1928ms | ✓ 1017ms | http |
| 121.126.185.63:25152 | 否 | 否 | ✓ 1664ms | ✓ 1984ms | ✓ 1832ms | http |
| 120.238.159.234:22222 | 否 | ✓ 1302ms | ✓ 1247ms | 否 | ✓ 971ms | http |
| 121.230.8.34:1080 | ✓ 1184ms | ✓ 1370ms | ✓ 1212ms | 否 | 否 | http |
| 106.117.208.101:7890 | ✓ 1695ms | 否 | ✓ 1613ms | 否 | ✓ 1648ms | http |
| 111.79.111.126:3128 | ✓ 1214ms | 否 | ✓ 1665ms | ✓ 1362ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1077ms | ✓ 1940ms | 否 | 否 | ✓ 1914ms | https |
| 180.103.19.53:1080 | 否 | ✓ 1555ms | ✓ 1125ms | ✓ 1498ms | ✓ 1132ms | http |
| 178.236.245.59:3128 | ✓ 1222ms | ✓ 1884ms | ✓ 1028ms | ✓ 1816ms | ✓ 1667ms | http |
| 178.236.245.17:3128 | ✓ 1235ms | 否 | ✓ 1065ms | ✓ 1986ms | ✓ 1376ms | http |
| 91.107.141.42:8081 | ✓ 1868ms | 否 | ✓ 1057ms | 否 | ✓ 1528ms | http |
| 162.248.165.72:1080 | ✓ 1612ms | 否 | ✓ 1906ms | 否 | ✓ 1953ms | http |
| 45.198.10.3:3125 | ✓ 1392ms | 否 | ✓ 1866ms | ✓ 1533ms | 否 | http |
| 35.181.173.74:41707 | ✓ 862ms | 否 | ✓ 1818ms | 否 | ✓ 1949ms | http |
| 212.175.29.184:8080 | ✓ 817ms | 否 | ✓ 1804ms | 否 | ✓ 1662ms | http |
| 103.39.51.190:8080 | ✓ 1918ms | 否 | 否 | ✓ 1511ms | ✓ 1476ms | http |
| 47.101.159.19:8899 | ✓ 994ms | ✓ 1139ms | ✓ 965ms | ✓ 1190ms | ✓ 940ms | http |
| 45.136.198.40:3128 | ✓ 1037ms | ✓ 1812ms | ✓ 1847ms | 否 | ✓ 1573ms | http |
| 45.186.6.104:3128 | ✓ 1590ms | ✓ 1683ms | ✓ 1657ms | 否 | 否 | http |
| 45.88.0.117:3128 | ✓ 1633ms | ✓ 1295ms | ✓ 1621ms | ✓ 1976ms | ✓ 1641ms | http |
| 202.129.206.239:3128 | ✓ 1217ms | 否 | 否 | ✓ 1647ms | ✓ 1315ms | http |
| 45.88.0.115:3128 | ✓ 497ms | ✓ 1604ms | ✓ 1355ms | ✓ 1979ms | ✓ 1603ms | http |
| 121.230.9.26:1080 | ✓ 1327ms | ✓ 1528ms | ✓ 1350ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 932ms | ✓ 1190ms | ✓ 980ms | ✓ 1255ms | ✓ 965ms | http |
| 62.113.119.14:8080 | ✓ 646ms | 否 | ✓ 990ms | ✓ 1603ms | ✓ 1127ms | http |
| 5.252.33.13:2025 | ✓ 1729ms | 否 | ✓ 1609ms | 否 | ✓ 1685ms | http |
| 45.88.0.116:3128 | ✓ 995ms | 否 | ✓ 794ms | 否 | ✓ 1424ms | http |
| 45.88.0.99:3128 | ✓ 997ms | ✓ 1486ms | ✓ 1316ms | 否 | ✓ 1456ms | http |
| 213.220.62.62:3128 | ✓ 992ms | ✓ 1789ms | ✓ 1018ms | 否 | ✓ 1457ms | http |
| 45.88.0.98:3128 | ✓ 996ms | ✓ 1656ms | ✓ 1139ms | 否 | ✓ 1463ms | http |
| 45.88.0.114:3128 | ✓ 1000ms | 否 | ✓ 798ms | 否 | ✓ 1463ms | http |
| 45.88.0.111:3128 | ✓ 997ms | 否 | ✓ 806ms | 否 | ✓ 1468ms | http |
| 45.88.0.113:3128 | ✓ 996ms | ✓ 1600ms | ✓ 1193ms | 否 | ✓ 1485ms | http |
| 104.248.81.109:3128 | ✓ 913ms | ✓ 1901ms | ✓ 1137ms | ✓ 1681ms | ✓ 1074ms | http |
| 18.201.114.187:7715 | ✓ 972ms | 否 | ✓ 1896ms | 否 | ✓ 1949ms | http |
| 116.80.96.106:3172 | ✓ 1625ms | 否 | ✓ 1731ms | 否 | ✓ 1696ms | http |
| 116.80.49.172:3172 | 否 | 否 | ✓ 1604ms | ✓ 1931ms | ✓ 1748ms | http |
| 109.234.38.35:3128 | ✓ 598ms | 否 | ✓ 1492ms | ✓ 1976ms | ✓ 1620ms | http |
| 45.140.147.155:1081 | ✓ 742ms | ✓ 1504ms | ✓ 1890ms | ✓ 1677ms | 否 | http |
| 172.212.68.37:3128 | 否 | ✓ 1554ms | ✓ 1932ms | 否 | ✓ 1792ms | http |
| 45.140.147.155:1082 | ✓ 710ms | 否 | ✓ 1390ms | ✓ 1689ms | 否 | http |

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
