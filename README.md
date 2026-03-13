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

最后更新：2026-03-13 03:26:17 UTC（2026-03-13 11:26:17 UTC+8）

**代理总数：81**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 431ms | ✓ 982ms | ✓ 1021ms | ✓ 1028ms | ✓ 671ms | http |
| 45.136.130.175:8443 | ✓ 479ms | ✓ 1038ms | ✓ 471ms | ✓ 902ms | ✓ 1703ms | http |
| 205.209.118.30:3138 | ✓ 401ms | ✓ 1276ms | ✓ 1224ms | ✓ 1122ms | ✓ 866ms | http |
| 1.231.81.166:3128 | ✓ 1698ms | 否 | ✓ 1105ms | ✓ 1140ms | ✓ 850ms | http |
| 113.160.132.26:8080 | ✓ 1557ms | ✓ 1504ms | ✓ 1422ms | ✓ 1436ms | ✓ 1433ms | http |
| 62.113.119.14:8080 | ✓ 1899ms | 否 | ✓ 1127ms | ✓ 1806ms | ✓ 1380ms | http |
| 152.42.213.210:8080 | ✓ 1707ms | 否 | ✓ 1669ms | ✓ 1558ms | ✓ 1027ms | http |
| 45.167.124.52:8080 | ✓ 1144ms | 否 | ✓ 1272ms | ✓ 1999ms | ✓ 1420ms | http |
| 185.115.74.185:8080 | ✓ 859ms | ✓ 1630ms | ✓ 1724ms | 否 | 否 | http |
| 103.113.70.189:1081 | ✓ 321ms | ✓ 1976ms | 否 | ✓ 1102ms | ✓ 911ms | http |
| 103.84.95.54:7890 | ✓ 799ms | 否 | ✓ 1125ms | 否 | ✓ 776ms | http |
| 186.148.180.46:999 | ✓ 1182ms | ✓ 1906ms | ✓ 1316ms | 否 | ✓ 1504ms | http |
| 165.227.5.10:8888 | ✓ 1095ms | ✓ 1787ms | ✓ 830ms | ✓ 1532ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1099ms | 否 | 否 | ✓ 1288ms | ✓ 1840ms | http |
| 59.46.216.131:30001 | ✓ 1044ms | 否 | ✓ 1247ms | ✓ 1570ms | 否 | http |
| 168.235.110.63:3128 | ✓ 435ms | 否 | ✓ 1152ms | ✓ 1356ms | 否 | http |
| 45.136.130.188:8443 | ✓ 426ms | ✓ 1034ms | ✓ 353ms | 否 | 否 | http |
| 121.230.8.144:1080 | 否 | ✓ 1674ms | 否 | ✓ 1464ms | ✓ 1976ms | http |
| 67.169.98.211:443 | ✓ 647ms | 否 | ✓ 1635ms | ✓ 1988ms | 否 | http |
| 86.53.183.16:1080 | ✓ 1610ms | 否 | ✓ 1483ms | 否 | ✓ 1439ms | http |
| 162.248.165.72:1080 | ✓ 1618ms | 否 | ✓ 1849ms | 否 | ✓ 1479ms | http |
| 81.70.169.194:80 | 否 | 否 | ✓ 1060ms | ✓ 1484ms | ✓ 1057ms | http |
| 101.43.255.96:80 | ✓ 1885ms | ✓ 1518ms | ✓ 1509ms | ✓ 1481ms | 否 | http |
| 190.9.109.198:999 | ✓ 522ms | 否 | ✓ 1267ms | ✓ 1498ms | ✓ 1126ms | http |
| 171.251.172.78:5104 | ✓ 1905ms | 否 | ✓ 1510ms | ✓ 1755ms | ✓ 1532ms | http |
| 46.183.25.8:443 | ✓ 832ms | 否 | ✓ 799ms | ✓ 1377ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1449ms | ✓ 1857ms | ✓ 1950ms | ✓ 1605ms | ✓ 1440ms | http |
| 193.168.173.136:443 | ✓ 744ms | ✓ 1743ms | ✓ 1990ms | 否 | ✓ 1840ms | http |
| 171.251.172.78:5107 | 否 | 否 | ✓ 1705ms | ✓ 1677ms | ✓ 1538ms | http |
| 14.225.212.37:7890 | ✓ 1282ms | ✓ 1694ms | ✓ 1307ms | 否 | ✓ 957ms | http |
| 45.136.131.63:8443 | ✓ 664ms | ✓ 895ms | ✓ 251ms | ✓ 909ms | ✓ 673ms | http |
| 45.136.130.191:8443 | ✓ 644ms | ✓ 855ms | ✓ 271ms | ✓ 1207ms | ✓ 668ms | http |
| 89.185.85.138:1080 | ✓ 685ms | 否 | ✓ 882ms | 否 | ✓ 1324ms | http |
| 185.241.5.57:3128 | 否 | 否 | ✓ 1041ms | ✓ 1915ms | ✓ 1586ms | http |
| 167.172.253.162:4857 | ✓ 1424ms | ✓ 1606ms | 否 | 否 | ✓ 986ms | http |
| 45.136.130.223:8443 | ✓ 243ms | ✓ 1316ms | ✓ 253ms | ✓ 878ms | ✓ 1693ms | http |
| 185.191.236.162:3128 | ✓ 913ms | ✓ 1568ms | ✓ 1735ms | 否 | ✓ 1092ms | http |
| 152.42.213.210:443 | ✓ 843ms | 否 | ✓ 1195ms | ✓ 1367ms | ✓ 997ms | http |
| 39.104.201.40:7890 | ✓ 1036ms | ✓ 1393ms | ✓ 1284ms | 否 | ✓ 1352ms | http |
| 150.107.140.238:3128 | ✓ 1544ms | 否 | ✓ 1578ms | ✓ 1590ms | ✓ 1008ms | http |
| 45.207.200.85:1080 | ✓ 1120ms | ✓ 1857ms | ✓ 971ms | 否 | 否 | http |
| 20.210.76.178:8561 | ✓ 1242ms | ✓ 1538ms | ✓ 712ms | ✓ 1163ms | ✓ 840ms | http |
| 157.254.37.238:999 | ✓ 876ms | 否 | ✓ 1343ms | ✓ 1440ms | ✓ 1243ms | http |
| 120.92.212.16:8890 | ✓ 1076ms | 否 | ✓ 1079ms | ✓ 1411ms | ✓ 1108ms | http |
| 120.92.212.16:7890 | ✓ 1063ms | 否 | 否 | ✓ 1369ms | ✓ 1105ms | http |
| 35.225.22.61:80 | ✓ 960ms | ✓ 1590ms | ✓ 1243ms | ✓ 840ms | ✓ 987ms | http |
| 45.88.0.99:3128 | ✓ 1167ms | ✓ 1429ms | ✓ 467ms | 否 | 否 | http |
| 45.88.0.114:3128 | ✓ 1794ms | ✓ 1421ms | ✓ 1510ms | 否 | ✓ 1195ms | http |
| 45.88.0.116:3128 | ✓ 1172ms | 否 | ✓ 827ms | 否 | ✓ 1922ms | http |
| 45.88.0.117:3128 | ✓ 1796ms | ✓ 1250ms | ✓ 1569ms | 否 | ✓ 1137ms | http |
| 34.101.184.164:3128 | ✓ 1834ms | 否 | ✓ 1600ms | ✓ 1749ms | ✓ 1123ms | http |
| 116.80.49.168:3172 | 否 | 否 | ✓ 1676ms | ✓ 1997ms | ✓ 1854ms | http |
| 45.88.0.115:3128 | ✓ 855ms | 否 | ✓ 950ms | ✓ 1334ms | ✓ 1060ms | http |
| 178.236.245.59:3128 | ✓ 1171ms | 否 | ✓ 1077ms | ✓ 1661ms | ✓ 1301ms | http |
| 121.230.9.148:1080 | ✓ 1370ms | ✓ 1614ms | ✓ 1239ms | ✓ 1665ms | ✓ 1180ms | http |
| 121.230.8.181:1080 | ✓ 1313ms | ✓ 1979ms | 否 | ✓ 1740ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1366ms | ✓ 1844ms | ✓ 1554ms | 否 | 否 | http |
| 13.36.243.194:9899 | ✓ 1066ms | 否 | ✓ 1455ms | 否 | ✓ 1581ms | http |
| 45.88.0.98:3128 | ✓ 1045ms | 否 | ✓ 1603ms | 否 | ✓ 1628ms | http |
| 137.184.1.87:3128 | ✓ 444ms | ✓ 1105ms | ✓ 1812ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 542ms | 否 | ✓ 1193ms | ✓ 1635ms | ✓ 1145ms | http |
| 88.80.150.82:8080 | ✓ 1201ms | ✓ 1876ms | 否 | 否 | ✓ 1778ms | https |
| 103.39.51.190:8080 | ✓ 1976ms | 否 | 否 | ✓ 1915ms | ✓ 1611ms | http |
| 180.103.19.219:1080 | 否 | 否 | ✓ 1341ms | ✓ 1786ms | ✓ 1274ms | http |
| 178.236.245.17:3128 | ✓ 1348ms | 否 | ✓ 1440ms | ✓ 1889ms | ✓ 1570ms | http |
| 45.136.198.40:3128 | ✓ 973ms | ✓ 1953ms | ✓ 1382ms | ✓ 1961ms | ✓ 1828ms | http |
| 192.71.213.85:9091 | ✓ 976ms | 否 | ✓ 1023ms | ✓ 1715ms | 否 | http |
| 37.187.109.70:10111 | ✓ 647ms | ✓ 1484ms | ✓ 1095ms | 否 | 否 | http |
| 34.141.27.50:3128 | ✓ 784ms | ✓ 1765ms | ✓ 1126ms | 否 | 否 | http |
| 103.82.93.219:3128 | ✓ 1030ms | 否 | ✓ 1204ms | ✓ 1364ms | ✓ 1401ms | http |
| 162.240.154.26:3128 | ✓ 1874ms | 否 | ✓ 1509ms | ✓ 1711ms | 否 | http |
| 121.230.8.213:1080 | ✓ 1165ms | 否 | 否 | ✓ 1792ms | ✓ 1109ms | http |
| 201.71.2.41:999 | ✓ 1005ms | ✓ 1616ms | ✓ 1535ms | ✓ 1644ms | ✓ 1458ms | http |
| 103.126.87.125:8090 | ✓ 1914ms | 否 | ✓ 1927ms | ✓ 1691ms | 否 | http |
| 103.139.138.194:3128 | ✓ 1880ms | 否 | ✓ 1652ms | ✓ 1716ms | 否 | http |
| 109.234.38.35:3128 | ✓ 1373ms | 否 | ✓ 1378ms | ✓ 1979ms | ✓ 1096ms | http |
| 106.117.208.101:7890 | ✓ 1714ms | 否 | ✓ 1967ms | 否 | ✓ 1735ms | http |
| 213.220.62.62:3128 | ✓ 1243ms | 否 | ✓ 1170ms | ✓ 1768ms | ✓ 1311ms | http |
| 45.88.0.113:3128 | ✓ 1244ms | 否 | ✓ 1166ms | ✓ 1737ms | ✓ 1340ms | http |
| 45.88.0.111:3128 | ✓ 1243ms | 否 | ✓ 1171ms | ✓ 1756ms | ✓ 1322ms | http |
| 45.140.147.155:1081 | ✓ 1231ms | 否 | ✓ 475ms | ✓ 1451ms | ✓ 1482ms | http |

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
