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

最后更新：2026-04-23 06:30:24 UTC（2026-04-23 14:30:24 UTC+8）

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
| 1.231.81.166:3128 | ✓ 906ms | ✓ 1211ms | ✓ 1165ms | ✓ 1113ms | ✓ 913ms | http |
| 34.96.238.40:8080 | ✓ 1367ms | 否 | ✓ 1296ms | ✓ 1529ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1316ms | 否 | ✓ 922ms | ✓ 1669ms | ✓ 1098ms | http |
| 218.108.131.186:17890 | ✓ 1510ms | ✓ 1751ms | ✓ 934ms | ✓ 1087ms | ✓ 980ms | http |
| 152.42.208.139:8118 | ✓ 1204ms | 否 | ✓ 1604ms | ✓ 1311ms | ✓ 1022ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1588ms | ✓ 1649ms | ✓ 1698ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 254ms | ✓ 885ms | ✓ 1513ms | http |
| 108.181.201.118:1234 | 否 | 否 | ✓ 976ms | ✓ 1011ms | ✓ 958ms | http |
| 59.46.216.131:30001 | ✓ 1281ms | 否 | ✓ 1271ms | ✓ 1604ms | ✓ 1725ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1399ms | ✓ 1355ms | 否 | ✓ 1143ms | http |
| 35.225.22.61:80 | ✓ 260ms | 否 | ✓ 1168ms | 否 | ✓ 962ms | http |
| 212.58.132.5:8888 | ✓ 1175ms | 否 | ✓ 1533ms | ✓ 1602ms | ✓ 1306ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 817ms | ✓ 1575ms | ✓ 1923ms | http |
| 120.92.212.16:8890 | ✓ 1096ms | ✓ 1617ms | ✓ 1270ms | 否 | 否 | http |
| 130.61.174.200:1080 | ✓ 656ms | ✓ 1703ms | ✓ 1921ms | ✓ 1925ms | ✓ 1403ms | http |
| 89.208.106.138:10808 | ✓ 1624ms | 否 | 否 | ✓ 1281ms | ✓ 903ms | http |
| 152.70.91.193:40000 | ✓ 1736ms | 否 | ✓ 1865ms | ✓ 1985ms | ✓ 1729ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1056ms | ✓ 1525ms | ✓ 1150ms | http |
| 128.199.121.61:9090 | ✓ 1004ms | 否 | ✓ 1064ms | ✓ 1590ms | ✓ 1821ms | http |
| 78.11.96.22:8888 | ✓ 823ms | ✓ 1463ms | ✓ 1648ms | ✓ 1466ms | ✓ 1281ms | http |
| 157.230.220.25:4857 | ✓ 880ms | ✓ 1972ms | 否 | ✓ 1315ms | 否 | http |
| 45.76.207.177:40000 | ✓ 1022ms | 否 | ✓ 1526ms | ✓ 1535ms | ✓ 1215ms | http |
| 45.140.147.155:1082 | ✓ 1017ms | ✓ 1298ms | 否 | 否 | ✓ 1008ms | http |
| 110.172.29.131:3128 | ✓ 1566ms | 否 | ✓ 1557ms | ✓ 1434ms | ✓ 1140ms | http |
| 223.84.151.86:30005 | ✓ 1804ms | ✓ 1720ms | ✓ 1264ms | ✓ 1780ms | ✓ 1609ms | http |
| 161.97.184.191:8080 | ✓ 1792ms | 否 | ✓ 1173ms | 否 | ✓ 1470ms | http |
| 34.71.229.255:3128 | ✓ 877ms | 否 | ✓ 1110ms | ✓ 1828ms | ✓ 1695ms | http |
| 162.240.154.26:3128 | ✓ 1976ms | 否 | ✓ 1449ms | ✓ 1460ms | ✓ 1835ms | http |
| 160.250.4.245:1 | ✓ 1873ms | 否 | 否 | ✓ 1649ms | ✓ 1521ms | http |
| 84.47.150.125:1080 | ✓ 1028ms | 否 | ✓ 1909ms | 否 | ✓ 1540ms | http |
| 195.26.224.49:3128 | ✓ 980ms | 否 | ✓ 740ms | ✓ 1613ms | 否 | http |
| 54.222.174.194:80 | 否 | ✓ 1953ms | ✓ 1815ms | ✓ 1966ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1689ms | 否 | ✓ 1167ms | 否 | ✓ 1845ms | http |
| 168.144.75.9:3128 | ✓ 1237ms | 否 | ✓ 1346ms | ✓ 1810ms | 否 | http |
| 177.93.132.244:3128 | ✓ 1362ms | 否 | ✓ 892ms | 否 | ✓ 1976ms | http |
| 45.140.147.82:1081 | ✓ 1140ms | ✓ 1835ms | ✓ 1827ms | 否 | 否 | http |
| 187.216.141.46:3128 | 否 | ✓ 1757ms | ✓ 1274ms | ✓ 1436ms | ✓ 1844ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1306ms | ✓ 1723ms | ✓ 1297ms | http |
| 45.59.122.132:80 | ✓ 448ms | 否 | ✓ 1517ms | 否 | ✓ 1035ms | http |
| 120.92.108.86:7890 | ✓ 1985ms | 否 | ✓ 1365ms | ✓ 1832ms | ✓ 1552ms | http |
| 183.232.248.73:7890 | ✓ 1024ms | ✓ 1369ms | ✓ 1086ms | ✓ 1298ms | ✓ 1073ms | http |
| 103.229.126.221:7890 | ✓ 1148ms | 否 | 否 | ✓ 1486ms | ✓ 1107ms | http |
| 121.230.8.89:1080 | ✓ 1222ms | 否 | ✓ 1252ms | ✓ 1941ms | ✓ 1340ms | http |
| 210.223.44.230:3128 | ✓ 1294ms | ✓ 1561ms | ✓ 1650ms | ✓ 1140ms | ✓ 1974ms | http |
| 45.134.14.169:8081 | ✓ 1407ms | 否 | ✓ 1587ms | ✓ 1896ms | ✓ 1636ms | http |
| 92.119.166.68:123 | ✓ 705ms | ✓ 1524ms | ✓ 1220ms | ✓ 1697ms | ✓ 1577ms | http |
| 114.237.77.245:1080 | 否 | ✓ 1358ms | ✓ 1180ms | ✓ 1506ms | ✓ 1205ms | http |
| 168.222.254.136:8888 | ✓ 1575ms | ✓ 1688ms | 否 | 否 | ✓ 1515ms | http |
| 8.219.195.129:1080 | ✓ 1203ms | 否 | ✓ 937ms | ✓ 1259ms | ✓ 1041ms | http |
| 85.190.99.143:443 | ✓ 791ms | ✓ 1926ms | ✓ 1497ms | ✓ 1740ms | ✓ 1301ms | http |
| 147.45.186.28:3128 | ✓ 872ms | ✓ 1753ms | ✓ 1316ms | ✓ 1614ms | ✓ 1285ms | http |
| 91.193.240.157:9877 | ✓ 1053ms | 否 | ✓ 1228ms | 否 | ✓ 1674ms | http |
| 91.99.15.45:2095 | ✓ 1053ms | ✓ 1946ms | ✓ 1914ms | 否 | 否 | http |
| 138.124.99.216:8888 | ✓ 1045ms | ✓ 1819ms | ✓ 1600ms | 否 | 否 | http |
| 217.77.102.18:3128 | ✓ 1620ms | 否 | ✓ 1922ms | 否 | ✓ 1917ms | http |
| 216.126.227.157:8080 | ✓ 858ms | ✓ 1475ms | ✓ 540ms | ✓ 1287ms | ✓ 1051ms | http |
| 84.47.150.126:1080 | ✓ 1471ms | ✓ 1791ms | ✓ 1742ms | ✓ 1993ms | ✓ 1921ms | http |
| 1.234.153.14:80 | ✓ 870ms | 否 | ✓ 923ms | ✓ 1078ms | ✓ 891ms | http |
| 150.107.140.238:3128 | ✓ 1834ms | 否 | 否 | ✓ 1417ms | ✓ 1069ms | http |
| 45.140.147.155:1081 | ✓ 1463ms | ✓ 1570ms | ✓ 1095ms | 否 | 否 | http |
| 83.219.250.8:62920 | ✓ 999ms | ✓ 1639ms | ✓ 1891ms | 否 | 否 | http |
| 113.176.92.71:3128 | ✓ 1888ms | 否 | ✓ 1624ms | ✓ 1911ms | 否 | http |
| 218.77.106.10:10150 | ✓ 1342ms | 否 | ✓ 1239ms | ✓ 1751ms | ✓ 1166ms | http |
| 183.98.143.134:8039 | ✓ 993ms | ✓ 1547ms | ✓ 984ms | ✓ 1252ms | ✓ 949ms | http |
| 47.84.73.61:1080 | ✓ 1020ms | ✓ 1922ms | ✓ 1036ms | ✓ 1280ms | ✓ 1018ms | http |
| 178.63.155.151:9002 | ✓ 919ms | 否 | ✓ 1472ms | ✓ 1983ms | ✓ 1584ms | http |
| 121.230.8.181:1080 | ✓ 1280ms | ✓ 1642ms | ✓ 1277ms | ✓ 1569ms | ✓ 1318ms | http |
| 117.236.124.166:3128 | ✓ 1659ms | 否 | ✓ 1825ms | ✓ 1868ms | ✓ 1633ms | http |
| 57.128.188.167:9165 | ✓ 1845ms | 否 | ✓ 1908ms | 否 | ✓ 1873ms | http |
| 103.129.200.2:8124 | ✓ 1533ms | 否 | ✓ 1367ms | 否 | ✓ 1368ms | http |
| 146.196.97.193:57413 | ✓ 1911ms | 否 | ✓ 1851ms | ✓ 1762ms | ✓ 1735ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 988ms | ✓ 1465ms | ✓ 1236ms | http |
| 213.220.3.234:20573 | ✓ 602ms | ✓ 1564ms | ✓ 1232ms | ✓ 1694ms | 否 | http |
| 101.255.157.6:8080 | 否 | 否 | ✓ 1534ms | ✓ 1958ms | ✓ 1710ms | http |
| 160.250.5.22:1 | ✓ 1800ms | 否 | 否 | ✓ 1594ms | ✓ 1216ms | http |

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
