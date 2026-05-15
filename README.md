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

最后更新：2026-05-15 13:47:17 UTC（2026-05-15 21:47:17 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1903ms | ✓ 1363ms | ✓ 1623ms | ✓ 1166ms | http |
| 129.80.217.21:444 | ✓ 1569ms | 否 | ✓ 808ms | ✓ 893ms | ✓ 643ms | http |
| 212.224.88.212:443 | ✓ 1274ms | 否 | ✓ 1219ms | ✓ 1890ms | ✓ 1462ms | http |
| 129.212.224.122:3128 | ✓ 1138ms | 否 | ✓ 1283ms | ✓ 1296ms | ✓ 1004ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1951ms | ✓ 1202ms | ✓ 1077ms | http |
| 137.59.47.73:3128 | ✓ 1367ms | 否 | ✓ 1741ms | ✓ 1437ms | ✓ 1102ms | http |
| 45.144.28.116:3128 | ✓ 1162ms | ✓ 1323ms | ✓ 776ms | 否 | 否 | http |
| 45.59.122.132:80 | ✓ 1617ms | ✓ 1813ms | ✓ 1639ms | ✓ 1475ms | ✓ 926ms | http |
| 160.238.65.9:3128 | ✓ 492ms | 否 | ✓ 474ms | ✓ 1374ms | 否 | http |
| 185.191.236.162:3128 | ✓ 675ms | 否 | ✓ 856ms | 否 | ✓ 1450ms | http |
| 45.88.0.99:3128 | ✓ 475ms | ✓ 1827ms | 否 | 否 | ✓ 963ms | http |
| 218.108.131.186:17890 | ✓ 1189ms | ✓ 1123ms | ✓ 938ms | ✓ 1138ms | ✓ 913ms | http |
| 212.58.132.5:8888 | ✓ 1829ms | 否 | ✓ 1051ms | ✓ 1578ms | ✓ 1277ms | http |
| 120.92.212.16:7890 | ✓ 967ms | 否 | 否 | ✓ 1359ms | ✓ 1336ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1565ms | 否 | ✓ 1804ms | ✓ 1166ms | http |
| 158.160.215.167:8127 | ✓ 1065ms | ✓ 1746ms | ✓ 584ms | ✓ 1825ms | ✓ 1193ms | http |
| 158.160.215.167:8123 | ✓ 1070ms | 否 | ✓ 654ms | 否 | ✓ 1333ms | http |
| 91.242.229.129:8092 | ✓ 1109ms | ✓ 1443ms | 否 | ✓ 1460ms | ✓ 1152ms | http |
| 158.160.215.167:8125 | ✓ 719ms | 否 | ✓ 1216ms | ✓ 1800ms | ✓ 1420ms | http |
| 34.71.229.255:3128 | ✓ 140ms | ✓ 1846ms | ✓ 159ms | ✓ 900ms | ✓ 690ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1332ms | ✓ 1285ms | ✓ 1042ms | http |
| 8.154.21.175:3128 | ✓ 871ms | ✓ 1090ms | ✓ 899ms | ✓ 1283ms | ✓ 922ms | http |
| 146.190.80.158:9090 | ✓ 975ms | 否 | ✓ 921ms | ✓ 1357ms | 否 | http |
| 128.199.121.61:9090 | ✓ 927ms | 否 | ✓ 1005ms | ✓ 1337ms | 否 | http |
| 128.199.116.219:9090 | ✓ 974ms | 否 | ✓ 956ms | ✓ 1338ms | ✓ 1102ms | http |
| 5.252.33.13:2025 | 否 | 否 | ✓ 1266ms | ✓ 1958ms | ✓ 1605ms | http |
| 160.238.65.6:3128 | ✓ 1908ms | 否 | ✓ 1981ms | ✓ 1348ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1123ms | ✓ 1811ms | ✓ 1816ms | ✓ 1860ms | ✓ 1443ms | http |
| 128.199.114.189:9090 | ✓ 932ms | 否 | ✓ 1553ms | ✓ 1958ms | 否 | http |
| 43.167.192.85:8080 | ✓ 1653ms | 否 | ✓ 804ms | ✓ 1074ms | ✓ 863ms | http |
| 168.110.52.228:3128 | ✓ 1650ms | ✓ 1729ms | 否 | 否 | ✓ 1201ms | http |
| 46.39.105.157:8080 | 否 | 否 | ✓ 1527ms | ✓ 1828ms | ✓ 1454ms | http |
| 210.223.44.230:3128 | ✓ 1997ms | ✓ 1301ms | ✓ 1295ms | ✓ 1176ms | ✓ 1298ms | http |
| 158.160.215.167:8124 | ✓ 1299ms | 否 | ✓ 1490ms | ✓ 1892ms | 否 | http |
| 8.219.97.248:80 | ✓ 1797ms | 否 | ✓ 1526ms | 否 | ✓ 1604ms | http |
| 20.164.75.153:8080 | ✓ 1602ms | 否 | ✓ 1605ms | 否 | ✓ 1970ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1599ms | ✓ 1972ms | ✓ 1425ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1003ms | 否 | 否 | ✓ 1250ms | ✓ 1000ms | http |
| 166.88.55.83:7890 | ✓ 825ms | ✓ 1368ms | ✓ 921ms | ✓ 1033ms | ✓ 871ms | http |
| 103.21.220.141:3128 | ✓ 861ms | 否 | ✓ 976ms | ✓ 1065ms | ✓ 826ms | http |
| 139.198.113.42:10023 | ✓ 892ms | ✓ 1461ms | ✓ 1409ms | 否 | ✓ 1340ms | http |
| 178.63.155.151:8888 | ✓ 1302ms | 否 | ✓ 1939ms | 否 | ✓ 1690ms | http |
| 20.78.118.91:8561 | ✓ 1565ms | ✓ 1447ms | ✓ 723ms | ✓ 1099ms | ✓ 866ms | http |
| 120.92.212.16:8890 | ✓ 1091ms | 否 | ✓ 1107ms | ✓ 1383ms | ✓ 1109ms | http |
| 104.248.151.93:9090 | ✓ 965ms | 否 | 否 | ✓ 1370ms | ✓ 1098ms | http |
| 20.27.13.35:8561 | ✓ 719ms | ✓ 1831ms | ✓ 818ms | ✓ 1567ms | ✓ 898ms | http |
| 159.89.31.62:8080 | ✓ 437ms | ✓ 1611ms | ✓ 1483ms | ✓ 1961ms | 否 | http |
| 128.199.113.85:9090 | ✓ 992ms | 否 | ✓ 991ms | ✓ 1364ms | ✓ 1122ms | http |
| 103.157.200.126:3128 | ✓ 1711ms | 否 | ✓ 1572ms | 否 | ✓ 1862ms | http |
| 160.238.65.3:3128 | ✓ 430ms | 否 | ✓ 1181ms | ✓ 1380ms | 否 | http |
| 3.101.133.120:80 | ✓ 506ms | ✓ 1342ms | ✓ 942ms | ✓ 1310ms | ✓ 1062ms | http |
| 192.142.18.232:8080 | ✓ 901ms | 否 | ✓ 508ms | ✓ 1961ms | ✓ 1765ms | http |
| 84.47.150.125:1080 | ✓ 1082ms | 否 | ✓ 486ms | 否 | ✓ 1604ms | http |
| 45.88.0.116:3128 | 否 | ✓ 1876ms | ✓ 1494ms | ✓ 1581ms | 否 | http |
| 20.210.39.153:8561 | ✓ 1578ms | ✓ 1096ms | ✓ 756ms | ✓ 1070ms | ✓ 882ms | http |
| 77.110.107.80:8080 | ✓ 1370ms | 否 | ✓ 1806ms | ✓ 1720ms | 否 | http |
| 103.147.152.12:1080 | ✓ 1214ms | 否 | 否 | ✓ 1448ms | ✓ 1169ms | http |
| 20.78.26.206:8561 | ✓ 762ms | 否 | ✓ 694ms | ✓ 1092ms | ✓ 810ms | http |
| 150.107.140.238:3128 | ✓ 1847ms | 否 | ✓ 1009ms | ✓ 1422ms | 否 | http |
| 77.110.107.80:1080 | ✓ 1424ms | ✓ 1791ms | ✓ 1245ms | ✓ 1819ms | 否 | http |
| 20.27.15.111:8561 | ✓ 1420ms | ✓ 1167ms | ✓ 927ms | ✓ 1187ms | ✓ 928ms | http |
| 128.199.254.13:9090 | ✓ 1550ms | 否 | 否 | ✓ 1402ms | ✓ 1190ms | http |
| 87.120.222.214:444 | ✓ 1186ms | 否 | 否 | ✓ 1972ms | ✓ 1579ms | http |
| 152.70.91.193:40000 | ✓ 1840ms | 否 | 否 | ✓ 1895ms | ✓ 1854ms | http |
| 91.217.81.131:1080 | ✓ 1321ms | 否 | ✓ 1288ms | 否 | ✓ 1502ms | http |
| 158.160.215.167:8126 | ✓ 1993ms | ✓ 1608ms | ✓ 665ms | ✓ 1789ms | ✓ 1575ms | http |
| 213.220.62.63:3128 | ✓ 1783ms | 否 | ✓ 1716ms | ✓ 1301ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1131ms | 否 | ✓ 1894ms | ✓ 1496ms | ✓ 1738ms | http |
| 57.129.144.178:40000 | ✓ 610ms | 否 | ✓ 907ms | ✓ 1653ms | ✓ 1545ms | http |
| 45.236.129.64:3128 | ✓ 1382ms | 否 | ✓ 1601ms | 否 | ✓ 1878ms | http |
| 103.235.67.190:80 | ✓ 1035ms | 否 | 否 | ✓ 1464ms | ✓ 1262ms | http |
| 61.52.131.172:8443 | ✓ 980ms | 否 | ✓ 1000ms | ✓ 1280ms | ✓ 1097ms | http |
| 160.238.65.2:3128 | ✓ 1153ms | ✓ 1883ms | ✓ 1188ms | 否 | 否 | http |
| 223.16.170.103:80 | ✓ 1603ms | 否 | ✓ 1775ms | ✓ 1552ms | 否 | http |
| 101.32.243.189:80 | 否 | ✓ 1877ms | ✓ 1850ms | 否 | ✓ 1559ms | http |
| 119.28.51.157:3128 | ✓ 1171ms | 否 | ✓ 1913ms | 否 | ✓ 1337ms | http |

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
