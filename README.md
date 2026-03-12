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

最后更新：2026-03-12 15:48:39 UTC（2026-03-12 23:48:39 UTC+8）

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
| 113.160.132.26:8080 | 否 | ✓ 1639ms | ✓ 1572ms | ✓ 1838ms | ✓ 1227ms | http |
| 1.231.81.166:3128 | ✓ 1723ms | ✓ 1678ms | ✓ 1408ms | ✓ 1419ms | 否 | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 808ms | ✓ 1275ms | ✓ 811ms | http |
| 91.107.141.42:8081 | ✓ 1000ms | 否 | ✓ 1880ms | ✓ 1853ms | 否 | http |
| 171.251.172.78:5106 | 否 | 否 | ✓ 1855ms | ✓ 1858ms | ✓ 1846ms | http |
| 62.113.119.14:8080 | ✓ 610ms | 否 | ✓ 813ms | ✓ 1595ms | ✓ 1473ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1147ms | ✓ 1444ms | ✓ 1198ms | http |
| 217.76.245.80:999 | ✓ 970ms | 否 | ✓ 1207ms | ✓ 1504ms | ✓ 1223ms | http |
| 190.9.109.198:999 | ✓ 718ms | 否 | ✓ 1148ms | ✓ 1320ms | ✓ 1221ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1468ms | ✓ 1757ms | 否 | ✓ 1152ms | http |
| 171.251.172.78:5104 | 否 | 否 | ✓ 1954ms | ✓ 1791ms | ✓ 1697ms | http |
| 171.251.172.78:5107 | 否 | 否 | ✓ 1751ms | ✓ 1859ms | ✓ 1736ms | http |
| 45.136.131.47:8443 | ✓ 664ms | ✓ 1252ms | ✓ 542ms | ✓ 906ms | ✓ 696ms | http |
| 113.59.32.162:22222 | ✓ 1276ms | ✓ 1567ms | ✓ 1168ms | 否 | 否 | http |
| 86.53.183.16:1080 | ✓ 1254ms | ✓ 1698ms | ✓ 1252ms | 否 | 否 | http |
| 202.155.12.161:443 | ✓ 1534ms | 否 | 否 | ✓ 1200ms | ✓ 1110ms | http |
| 14.225.212.37:7890 | ✓ 1645ms | 否 | 否 | ✓ 1583ms | ✓ 1031ms | http |
| 45.136.130.191:8443 | ✓ 531ms | ✓ 924ms | ✓ 275ms | ✓ 908ms | ✓ 699ms | http |
| 45.136.130.188:8443 | ✓ 525ms | ✓ 1910ms | ✓ 272ms | ✓ 899ms | ✓ 678ms | http |
| 45.136.130.175:8443 | ✓ 530ms | ✓ 1420ms | ✓ 267ms | ✓ 903ms | ✓ 1290ms | http |
| 152.42.213.210:8080 | ✓ 888ms | 否 | 否 | ✓ 1495ms | ✓ 1023ms | http |
| 120.240.35.178:22222 | ✓ 1150ms | 否 | ✓ 1450ms | ✓ 1536ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1939ms | ✓ 1855ms | ✓ 1652ms | 否 | 否 | http |
| 45.136.131.63:8443 | ✓ 586ms | ✓ 1373ms | ✓ 301ms | ✓ 953ms | ✓ 706ms | http |
| 117.159.239.51:22222 | ✓ 984ms | 否 | ✓ 980ms | ✓ 1666ms | 否 | http |
| 113.59.32.142:22222 | ✓ 1239ms | ✓ 1548ms | ✓ 1111ms | ✓ 1420ms | ✓ 1182ms | http |
| 46.183.25.8:443 | ✓ 1415ms | 否 | ✓ 814ms | ✓ 1537ms | 否 | http |
| 120.240.35.177:22222 | 否 | 否 | ✓ 1157ms | ✓ 1527ms | ✓ 1535ms | http |
| 222.184.48.252:22222 | ✓ 1785ms | ✓ 1314ms | ✓ 1327ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1092ms | ✓ 411ms | ✓ 1129ms | ✓ 893ms | http |
| 183.249.5.117:22222 | ✓ 930ms | 否 | ✓ 931ms | ✓ 1203ms | ✓ 937ms | http |
| 117.159.239.44:22222 | ✓ 1044ms | ✓ 1244ms | ✓ 984ms | 否 | 否 | http |
| 120.238.159.189:22222 | 否 | ✓ 1421ms | ✓ 1143ms | 否 | ✓ 1111ms | http |
| 120.240.35.173:22222 | ✓ 1169ms | ✓ 1751ms | ✓ 1192ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 1760ms | 否 | 否 | ✓ 1525ms | ✓ 986ms | http |
| 196.204.141.249:1976 | 否 | 否 | ✓ 1145ms | ✓ 1721ms | ✓ 1491ms | http |
| 222.184.48.248:22222 | ✓ 1381ms | ✓ 1641ms | ✓ 1024ms | 否 | ✓ 1105ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1636ms | 否 | ✓ 1106ms | ✓ 806ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1585ms | ✓ 481ms | ✓ 1363ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1520ms | ✓ 1718ms | ✓ 854ms | ✓ 1730ms | ✓ 1364ms | http |
| 168.235.110.63:3128 | ✓ 1163ms | 否 | ✓ 1734ms | ✓ 1652ms | 否 | http |
| 120.238.159.230:22222 | 否 | ✓ 1564ms | ✓ 1504ms | ✓ 1493ms | 否 | http |
| 121.230.9.198:1080 | ✓ 1876ms | ✓ 1716ms | ✓ 1400ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1938ms | ✓ 1705ms | 否 | 否 | ✓ 1177ms | http |
| 178.236.245.17:3128 | 否 | 否 | ✓ 1171ms | ✓ 1822ms | ✓ 1788ms | http |
| 178.236.245.59:3128 | 否 | 否 | ✓ 1131ms | ✓ 1876ms | ✓ 1761ms | http |
| 39.104.201.40:7890 | ✓ 1660ms | ✓ 1647ms | ✓ 1058ms | ✓ 1471ms | ✓ 1394ms | http |
| 45.136.198.40:3128 | ✓ 982ms | ✓ 1822ms | 否 | 否 | ✓ 1565ms | http |
| 45.136.130.223:8443 | ✓ 1377ms | 否 | ✓ 1417ms | ✓ 1340ms | ✓ 1141ms | http |
| 183.249.5.213:22222 | ✓ 891ms | ✓ 1054ms | ✓ 981ms | ✓ 1168ms | 否 | http |
| 117.159.239.54:22222 | 否 | 否 | ✓ 1253ms | ✓ 1564ms | ✓ 1004ms | http |
| 120.240.35.175:22222 | ✓ 1183ms | 否 | ✓ 1242ms | 否 | ✓ 1056ms | http |
| 120.198.141.79:22222 | ✓ 1119ms | ✓ 1571ms | ✓ 1160ms | ✓ 1458ms | ✓ 1082ms | http |
| 115.231.181.40:8128 | ✓ 1429ms | ✓ 1333ms | ✓ 1022ms | 否 | ✓ 1144ms | http |
| 113.59.32.161:22222 | ✓ 1220ms | 否 | 否 | ✓ 1484ms | ✓ 1107ms | http |
| 101.32.244.83:8080 | ✓ 1186ms | 否 | ✓ 1157ms | ✓ 1640ms | ✓ 1475ms | http |
| 121.43.196.213:8222 | ✓ 1113ms | ✓ 1274ms | ✓ 1039ms | ✓ 1320ms | ✓ 1055ms | http |
| 121.43.196.210:8222 | ✓ 1121ms | ✓ 1278ms | ✓ 1032ms | ✓ 1387ms | ✓ 1046ms | http |
| 114.55.226.123:10086 | ✓ 1191ms | ✓ 1562ms | ✓ 1216ms | ✓ 1487ms | ✓ 1355ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1279ms | ✓ 1654ms | 否 | ✓ 1051ms | http |
| 81.70.169.194:80 | 否 | ✓ 1552ms | 否 | ✓ 1768ms | ✓ 1718ms | http |
| 119.18.145.49:20326 | ✓ 1895ms | 否 | ✓ 1570ms | ✓ 1885ms | 否 | http |
| 120.198.141.75:22222 | ✓ 1163ms | ✓ 1476ms | ✓ 1219ms | ✓ 1439ms | ✓ 1143ms | http |
| 120.240.29.51:22222 | ✓ 1277ms | 否 | 否 | ✓ 1770ms | ✓ 1227ms | http |
| 162.248.165.72:1080 | ✓ 1045ms | ✓ 1977ms | ✓ 613ms | 否 | 否 | http |
| 124.16.111.161:7890 | ✓ 1300ms | ✓ 1797ms | ✓ 1140ms | ✓ 1371ms | ✓ 1057ms | http |
| 178.156.224.42:3128 | ✓ 1023ms | 否 | 否 | ✓ 1921ms | ✓ 1489ms | http |
| 120.240.35.160:22222 | ✓ 1067ms | 否 | 否 | ✓ 1446ms | ✓ 1090ms | http |
| 107.173.52.58:7890 | ✓ 1119ms | 否 | ✓ 1166ms | 否 | ✓ 877ms | http |
| 61.0.226.241:3128 | ✓ 1427ms | 否 | ✓ 1690ms | 否 | ✓ 1633ms | http |
| 167.172.67.118:8080 | ✓ 1005ms | 否 | ✓ 1584ms | 否 | ✓ 1013ms | http |
| 93.174.125.63:80 | ✓ 1839ms | 否 | 否 | ✓ 1462ms | ✓ 1157ms | http |
| 183.249.5.214:22222 | ✓ 863ms | ✓ 1052ms | ✓ 925ms | ✓ 1178ms | ✓ 885ms | http |
| 120.240.35.176:22222 | ✓ 1204ms | ✓ 1447ms | 否 | 否 | ✓ 1113ms | http |
| 150.241.77.125:3128 | ✓ 897ms | 否 | ✓ 1445ms | ✓ 1886ms | ✓ 1328ms | http |
| 45.236.129.64:3128 | ✓ 1721ms | ✓ 1978ms | ✓ 1744ms | 否 | ✓ 1691ms | http |
| 222.184.48.241:22222 | ✓ 1816ms | 否 | ✓ 1470ms | 否 | ✓ 1330ms | http |
| 103.39.51.190:8080 | ✓ 1822ms | 否 | ✓ 1589ms | 否 | ✓ 1591ms | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1939ms | ✓ 1963ms | ✓ 1908ms | http |
| 183.249.5.109:22222 | ✓ 1024ms | ✓ 1095ms | ✓ 940ms | ✓ 1399ms | ✓ 1006ms | http |
| 45.88.0.115:3128 | ✓ 556ms | 否 | ✓ 1566ms | 否 | ✓ 1730ms | http |
| 45.88.0.117:3128 | ✓ 1649ms | 否 | ✓ 1566ms | ✓ 1834ms | ✓ 1901ms | http |
| 183.249.5.111:22222 | ✓ 873ms | ✓ 1329ms | ✓ 967ms | ✓ 1182ms | ✓ 891ms | http |
| 45.168.238.193:8443 | ✓ 750ms | ✓ 1147ms | ✓ 318ms | ✓ 1127ms | 否 | http |

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
