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

最后更新：2026-04-19 14:03:09 UTC（2026-04-19 22:03:09 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:3128 | ✓ 990ms | ✓ 1645ms | 否 | ✓ 1465ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1543ms | ✓ 1433ms | ✓ 1356ms | ✓ 1484ms | ✓ 995ms | http |
| 14.247.76.52:8080 | ✓ 1519ms | 否 | ✓ 919ms | ✓ 1664ms | ✓ 1071ms | http |
| 188.246.224.49:7890 | ✓ 687ms | ✓ 1734ms | 否 | 否 | ✓ 1635ms | http |
| 46.101.95.183:8888 | ✓ 930ms | 否 | 否 | ✓ 1823ms | ✓ 1973ms | http |
| 152.42.208.139:8118 | 否 | 否 | ✓ 1370ms | ✓ 1184ms | ✓ 937ms | http |
| 1.231.81.166:3128 | ✓ 1387ms | 否 | ✓ 1876ms | ✓ 1053ms | ✓ 796ms | http |
| 149.51.42.10:8080 | ✓ 674ms | ✓ 1307ms | 否 | ✓ 1718ms | 否 | http |
| 47.85.84.192:20000 | ✓ 266ms | 否 | ✓ 563ms | ✓ 1396ms | ✓ 1190ms | http |
| 177.93.132.244:3128 | ✓ 742ms | 否 | ✓ 855ms | 否 | ✓ 1715ms | http |
| 159.89.191.221:3128 | ✓ 400ms | 否 | 否 | ✓ 1439ms | ✓ 1544ms | http |
| 38.55.106.208:6005 | ✓ 945ms | 否 | ✓ 1181ms | ✓ 1727ms | ✓ 879ms | http |
| 194.104.9.38:3128 | ✓ 1177ms | 否 | ✓ 1573ms | ✓ 1397ms | ✓ 1689ms | http |
| 223.84.151.86:30005 | ✓ 1592ms | ✓ 1348ms | ✓ 1460ms | ✓ 1638ms | ✓ 1704ms | http |
| 115.231.181.40:8128 | ✓ 1033ms | ✓ 1266ms | ✓ 960ms | 否 | 否 | http |
| 159.223.225.118:8888 | ✓ 1137ms | 否 | 否 | ✓ 1422ms | ✓ 1398ms | http |
| 78.11.96.22:8888 | ✓ 1365ms | 否 | ✓ 1082ms | ✓ 1906ms | ✓ 1374ms | http |
| 185.138.116.150:8080 | ✓ 648ms | ✓ 1794ms | ✓ 760ms | ✓ 1979ms | ✓ 1412ms | http |
| 91.99.15.45:2095 | ✓ 660ms | 否 | ✓ 1114ms | 否 | ✓ 1424ms | http |
| 152.32.132.190:7890 | ✓ 1058ms | 否 | ✓ 1630ms | ✓ 1033ms | ✓ 1649ms | http |
| 162.19.253.202:8443 | ✓ 802ms | 否 | ✓ 1979ms | 否 | ✓ 1979ms | http |
| 121.135.144.234:8972 | 否 | 否 | ✓ 1651ms | ✓ 1637ms | ✓ 1150ms | http |
| 120.92.212.16:8890 | ✓ 975ms | 否 | ✓ 1100ms | 否 | ✓ 1328ms | http |
| 120.92.212.16:7890 | ✓ 1021ms | ✓ 1366ms | ✓ 1065ms | 否 | 否 | http |
| 166.88.61.54:8000 | ✓ 1324ms | ✓ 1597ms | ✓ 1114ms | ✓ 961ms | ✓ 1807ms | http |
| 20.78.26.206:8561 | ✓ 1624ms | 否 | ✓ 767ms | ✓ 901ms | 否 | http |
| 85.190.99.143:443 | ✓ 682ms | 否 | ✓ 1549ms | 否 | ✓ 1724ms | http |
| 116.171.106.111:3443 | ✓ 1522ms | ✓ 1824ms | ✓ 1971ms | 否 | 否 | http |
| 20.27.14.220:8561 | ✓ 1018ms | ✓ 1259ms | ✓ 735ms | ✓ 909ms | ✓ 786ms | http |
| 20.27.15.111:8561 | ✓ 1020ms | ✓ 1218ms | ✓ 779ms | ✓ 916ms | ✓ 776ms | http |
| 62.113.119.14:8080 | ✓ 1232ms | 否 | ✓ 687ms | ✓ 1953ms | ✓ 1200ms | http |
| 45.140.147.82:1082 | ✓ 1199ms | ✓ 1871ms | ✓ 1622ms | ✓ 1681ms | ✓ 1208ms | http |
| 84.47.150.125:1080 | ✓ 1232ms | 否 | ✓ 1769ms | 否 | ✓ 1671ms | http |
| 59.46.216.131:30001 | ✓ 1039ms | ✓ 1543ms | ✓ 1719ms | ✓ 1440ms | 否 | http |
| 36.141.21.200:7890 | 否 | ✓ 1710ms | ✓ 1137ms | ✓ 1339ms | ✓ 1077ms | http |
| 45.12.151.226:2829 | ✓ 696ms | 否 | ✓ 1105ms | 否 | ✓ 1714ms | http |
| 20.210.39.153:8561 | ✓ 678ms | ✓ 974ms | ✓ 853ms | ✓ 998ms | ✓ 796ms | http |
| 168.144.75.9:3128 | ✓ 1725ms | 否 | ✓ 1061ms | ✓ 1834ms | 否 | http |
| 210.223.44.230:3128 | ✓ 720ms | ✓ 1486ms | ✓ 1284ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1318ms | 否 | ✓ 1734ms | ✓ 1902ms | ✓ 1851ms | http |
| 102.134.49.165:6005 | ✓ 619ms | 否 | ✓ 1652ms | ✓ 1113ms | ✓ 804ms | http |
| 47.95.231.180:8084 | ✓ 943ms | ✓ 1341ms | ✓ 1063ms | ✓ 1281ms | ✓ 1017ms | http |
| 102.134.48.240:6005 | ✓ 1743ms | ✓ 1579ms | ✓ 1834ms | ✓ 1030ms | ✓ 1014ms | http |
| 38.55.104.8:6005 | 否 | 否 | ✓ 1148ms | ✓ 1083ms | ✓ 836ms | http |
| 38.55.106.206:6005 | 否 | 否 | ✓ 1064ms | ✓ 1253ms | ✓ 857ms | http |
| 38.55.104.182:6005 | 否 | 否 | ✓ 1063ms | ✓ 1135ms | ✓ 1025ms | http |
| 38.55.107.254:6005 | 否 | 否 | ✓ 1248ms | ✓ 1059ms | ✓ 875ms | http |
| 38.55.107.137:6005 | 否 | 否 | ✓ 1922ms | ✓ 1345ms | ✓ 854ms | http |
| 113.192.1.138:8181 | ✓ 1950ms | 否 | ✓ 1365ms | 否 | ✓ 1487ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1212ms | 否 | ✓ 971ms | ✓ 787ms | http |
| 168.110.52.228:3128 | ✓ 544ms | 否 | ✓ 810ms | ✓ 1436ms | ✓ 1323ms | http |
| 20.120.225.109:3128 | ✓ 720ms | ✓ 1291ms | ✓ 918ms | ✓ 1383ms | ✓ 1004ms | http |
| 45.93.29.147:6005 | ✓ 1273ms | ✓ 1325ms | 否 | ✓ 1647ms | ✓ 1889ms | http |
| 161.97.184.191:8080 | ✓ 1330ms | 否 | ✓ 1448ms | 否 | ✓ 1786ms | http |
| 35.225.22.61:80 | 否 | ✓ 1420ms | ✓ 1180ms | ✓ 1496ms | ✓ 931ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1799ms | 否 | ✓ 1887ms | ✓ 1432ms | http |
| 192.3.248.190:8014 | 否 | ✓ 1712ms | ✓ 1388ms | 否 | ✓ 1123ms | http |
| 5.63.111.238:8080 | ✓ 1862ms | 否 | ✓ 1987ms | 否 | ✓ 1885ms | http |
| 94.131.118.129:1081 | ✓ 1074ms | ✓ 1610ms | ✓ 727ms | ✓ 1949ms | ✓ 1503ms | http |
| 94.131.118.129:1082 | ✓ 811ms | 否 | ✓ 544ms | ✓ 1641ms | ✓ 1642ms | http |
| 208.87.243.199:7878 | 否 | ✓ 905ms | ✓ 906ms | ✓ 1013ms | ✓ 658ms | http |
| 72.56.105.251:3128 | 否 | 否 | ✓ 1084ms | ✓ 1554ms | ✓ 1447ms | http |
| 117.122.240.82:3338 | ✓ 1854ms | ✓ 1855ms | ✓ 1265ms | ✓ 1718ms | ✓ 1780ms | http |
| 81.30.156.115:8080 | ✓ 932ms | 否 | ✓ 883ms | ✓ 1847ms | ✓ 1576ms | http |
| 91.107.124.215:3128 | ✓ 855ms | ✓ 1715ms | ✓ 872ms | 否 | ✓ 1794ms | http |
| 103.18.77.14:1111 | ✓ 1529ms | 否 | 否 | ✓ 1907ms | ✓ 1603ms | http |
| 45.76.207.177:40000 | ✓ 799ms | 否 | ✓ 1162ms | ✓ 1115ms | ✓ 1013ms | http |
| 20.27.11.248:8561 | ✓ 573ms | ✓ 959ms | ✓ 647ms | ✓ 1019ms | ✓ 720ms | http |
| 20.27.13.35:8561 | ✓ 570ms | ✓ 1829ms | ✓ 589ms | ✓ 902ms | ✓ 690ms | http |
| 103.113.70.189:1081 | ✓ 1299ms | ✓ 1371ms | ✓ 1144ms | 否 | ✓ 950ms | http |
| 114.237.77.196:1080 | ✓ 1718ms | ✓ 1364ms | ✓ 1271ms | ✓ 1407ms | ✓ 1140ms | http |
| 77.110.113.24:40000 | ✓ 1323ms | 否 | ✓ 1500ms | 否 | ✓ 1735ms | http |
| 101.32.243.189:80 | ✓ 1274ms | 否 | ✓ 1630ms | ✓ 1333ms | ✓ 1330ms | http |
| 54.37.72.89:80 | ✓ 1256ms | ✓ 1721ms | ✓ 1993ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 952ms | ✓ 1298ms | ✓ 1026ms | ✓ 1365ms | ✓ 1039ms | http |
| 45.88.0.117:3128 | ✓ 516ms | ✓ 1274ms | ✓ 1713ms | ✓ 1932ms | ✓ 1567ms | http |
| 213.220.62.62:3128 | ✓ 560ms | ✓ 1393ms | ✓ 648ms | ✓ 1370ms | ✓ 1010ms | http |
| 103.113.70.189:1082 | ✓ 529ms | 否 | ✓ 213ms | 否 | ✓ 1081ms | http |
| 103.231.75.209:3128 | ✓ 1877ms | 否 | ✓ 1670ms | ✓ 1915ms | ✓ 1695ms | http |
| 43.132.188.134:443 | ✓ 754ms | ✓ 1571ms | 否 | 否 | ✓ 1585ms | http |

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
