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

最后更新：2026-03-10 00:24:42 UTC（2026-03-10 08:24:42 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 383ms | ✓ 794ms | ✓ 586ms | ✓ 672ms | ✓ 1539ms | http |
| 154.3.236.202:3128 | ✓ 516ms | ✓ 1509ms | ✓ 1229ms | ✓ 1546ms | ✓ 1113ms | http |
| 162.240.154.26:3128 | ✓ 907ms | ✓ 1622ms | ✓ 1271ms | ✓ 1264ms | ✓ 969ms | http |
| 61.72.221.94:3128 | ✓ 1149ms | 否 | ✓ 1144ms | ✓ 1004ms | ✓ 729ms | http |
| 61.72.110.114:3128 | 否 | 否 | ✓ 1374ms | ✓ 1018ms | ✓ 788ms | http |
| 202.155.12.161:443 | 否 | ✓ 1985ms | ✓ 1259ms | ✓ 1425ms | ✓ 1027ms | http |
| 101.47.73.135:3128 | ✓ 1439ms | 否 | 否 | ✓ 1559ms | ✓ 984ms | http |
| 1.231.81.166:3128 | ✓ 1413ms | ✓ 1208ms | ✓ 1683ms | ✓ 1346ms | ✓ 1235ms | http |
| 185.191.236.162:3128 | ✓ 1245ms | 否 | ✓ 1873ms | 否 | ✓ 1705ms | http |
| 95.3.9.78:3128 | ✓ 1532ms | 否 | 否 | ✓ 1759ms | ✓ 1351ms | http |
| 8.219.97.248:80 | ✓ 1862ms | 否 | ✓ 1248ms | 否 | ✓ 1949ms | http |
| 47.101.149.27:9010 | ✓ 1709ms | ✓ 1317ms | ✓ 1718ms | 否 | 否 | http |
| 61.72.110.54:3128 | ✓ 952ms | ✓ 1305ms | ✓ 990ms | 否 | ✓ 851ms | http |
| 159.223.42.219:3128 | ✓ 1333ms | 否 | ✓ 1467ms | ✓ 1221ms | ✓ 940ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1259ms | ✓ 991ms | ✓ 1206ms | 否 | http |
| 217.76.245.80:999 | ✓ 914ms | 否 | ✓ 1188ms | ✓ 1368ms | ✓ 1161ms | http |
| 115.231.181.40:8128 | ✓ 880ms | 否 | ✓ 1023ms | ✓ 1219ms | ✓ 985ms | http |
| 181.63.102.155:8080 | ✓ 1251ms | ✓ 1725ms | ✓ 1562ms | 否 | ✓ 1636ms | http |
| 35.225.22.61:80 | ✓ 425ms | ✓ 1265ms | ✓ 572ms | ✓ 1210ms | 否 | http |
| 165.227.5.10:8888 | ✓ 1319ms | 否 | 否 | ✓ 743ms | ✓ 722ms | http |
| 91.107.141.42:8081 | ✓ 1236ms | ✓ 1985ms | ✓ 1529ms | ✓ 1675ms | ✓ 1593ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1564ms | ✓ 894ms | 否 | ✓ 1636ms | http |
| 116.58.162.45:3128 | ✓ 1179ms | ✓ 1560ms | ✓ 1181ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1311ms | 否 | ✓ 918ms | ✓ 1215ms | 否 | http |
| 101.43.255.96:80 | ✓ 876ms | ✓ 1300ms | ✓ 1083ms | ✓ 1191ms | ✓ 979ms | http |
| 81.70.169.194:80 | ✓ 1017ms | ✓ 1506ms | ✓ 1081ms | ✓ 1302ms | ✓ 986ms | http |
| 39.104.201.40:7890 | ✓ 1198ms | ✓ 1302ms | 否 | ✓ 1500ms | ✓ 1014ms | http |
| 152.42.213.210:8080 | ✓ 1669ms | 否 | ✓ 1484ms | ✓ 1118ms | 否 | http |
| 121.237.181.137:8888 | ✓ 898ms | ✓ 1276ms | ✓ 916ms | 否 | 否 | http |
| 190.9.109.198:999 | ✓ 710ms | ✓ 1478ms | ✓ 1461ms | ✓ 1659ms | ✓ 1222ms | http |
| 190.9.109.207:999 | ✓ 718ms | ✓ 1463ms | ✓ 1466ms | ✓ 1601ms | 否 | http |
| 46.183.25.8:443 | ✓ 1481ms | 否 | 否 | ✓ 1467ms | ✓ 1701ms | http |
| 193.168.173.136:443 | ✓ 1084ms | 否 | 否 | ✓ 1869ms | ✓ 1718ms | http |
| 194.213.18.200:443 | ✓ 1814ms | 否 | ✓ 1050ms | ✓ 1282ms | ✓ 1079ms | http |
| 107.152.32.98:1305 | ✓ 1182ms | 否 | 否 | ✓ 1965ms | ✓ 1678ms | http |
| 185.115.74.185:8080 | ✓ 969ms | ✓ 1887ms | ✓ 1551ms | 否 | 否 | http |
| 14.225.222.164:7890 | ✓ 1844ms | ✓ 1911ms | 否 | ✓ 1644ms | ✓ 886ms | http |
| 62.234.206.73:3128 | ✓ 1044ms | ✓ 1291ms | ✓ 863ms | ✓ 1154ms | 否 | http |
| 103.236.89.228:7890 | ✓ 1010ms | ✓ 1280ms | ✓ 1117ms | ✓ 1532ms | ✓ 1395ms | http |
| 114.55.226.123:10086 | ✓ 1074ms | ✓ 1364ms | ✓ 998ms | ✓ 1323ms | ✓ 1002ms | http |
| 45.140.147.155:1081 | ✓ 586ms | ✓ 1490ms | ✓ 1389ms | ✓ 1702ms | ✓ 1228ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1275ms | 否 | ✓ 1073ms | ✓ 1104ms | http |
| 178.236.245.59:3128 | ✓ 813ms | 否 | ✓ 1097ms | ✓ 1878ms | ✓ 1537ms | http |
| 94.176.3.43:7443 | ✓ 1054ms | ✓ 1783ms | ✓ 1831ms | 否 | ✓ 1586ms | http |
| 136.49.34.18:8888 | ✓ 364ms | 否 | ✓ 420ms | 否 | ✓ 855ms | http |
| 116.80.82.218:3172 | ✓ 1863ms | 否 | ✓ 1999ms | 否 | ✓ 1849ms | http |
| 45.136.130.223:8443 | 否 | ✓ 1241ms | ✓ 1200ms | ✓ 668ms | ✓ 622ms | http |
| 103.183.10.203:3125 | ✓ 1269ms | 否 | ✓ 1277ms | ✓ 1535ms | ✓ 1386ms | http |
| 101.32.244.83:8080 | ✓ 1100ms | ✓ 1741ms | ✓ 1213ms | ✓ 1378ms | ✓ 1231ms | http |
| 121.43.196.213:8222 | ✓ 955ms | ✓ 1082ms | ✓ 855ms | ✓ 1145ms | ✓ 863ms | http |
| 121.43.196.210:8222 | ✓ 915ms | ✓ 1111ms | ✓ 865ms | ✓ 1168ms | ✓ 891ms | http |
| 88.80.150.82:8080 | ✓ 1225ms | 否 | ✓ 1995ms | 否 | ✓ 1882ms | https |
| 162.248.165.72:1080 | ✓ 1514ms | 否 | ✓ 1953ms | 否 | ✓ 1828ms | http |
| 223.204.176.85:3128 | ✓ 1202ms | 否 | ✓ 1389ms | ✓ 1859ms | ✓ 1562ms | http |
| 210.223.44.230:3128 | ✓ 1090ms | 否 | 否 | ✓ 981ms | ✓ 1403ms | http |
| 201.144.20.238:3128 | 否 | 否 | ✓ 920ms | ✓ 1278ms | ✓ 1214ms | http |
| 45.186.6.104:3128 | ✓ 1293ms | ✓ 1741ms | ✓ 1788ms | 否 | 否 | http |
| 220.170.182.39:9293 | ✓ 1373ms | ✓ 1459ms | ✓ 1380ms | ✓ 1401ms | ✓ 1433ms | http |
| 103.113.70.189:1081 | ✓ 406ms | ✓ 1351ms | 否 | ✓ 1258ms | ✓ 981ms | http |
| 178.236.245.17:3128 | 否 | 否 | ✓ 872ms | ✓ 1761ms | ✓ 1626ms | http |
| 103.239.41.117:8080 | ✓ 1896ms | 否 | ✓ 1346ms | 否 | ✓ 1383ms | http |
| 103.39.51.190:8080 | ✓ 1874ms | 否 | 否 | ✓ 1533ms | ✓ 1531ms | http |
| 106.14.203.63:3333 | ✓ 960ms | ✓ 1505ms | ✓ 1766ms | 否 | ✓ 1633ms | http |
| 62.113.119.14:8080 | ✓ 1243ms | 否 | ✓ 1428ms | ✓ 1625ms | ✓ 1168ms | http |
| 45.136.198.40:3128 | ✓ 1232ms | ✓ 1620ms | ✓ 1323ms | 否 | ✓ 1966ms | http |
| 172.212.68.37:3128 | ✓ 1649ms | ✓ 1733ms | 否 | 否 | ✓ 1293ms | http |
| 168.235.110.63:3128 | ✓ 427ms | ✓ 1130ms | ✓ 304ms | ✓ 1240ms | ✓ 1045ms | http |
| 152.70.98.46:8888 | ✓ 674ms | 否 | ✓ 1522ms | ✓ 1038ms | ✓ 853ms | http |
| 103.67.46.225:3125 | ✓ 1697ms | 否 | 否 | ✓ 1735ms | ✓ 1555ms | http |
| 61.52.131.172:8443 | ✓ 855ms | ✓ 1131ms | ✓ 945ms | ✓ 1107ms | ✓ 971ms | http |
| 5.252.33.13:2025 | ✓ 1756ms | 否 | ✓ 1383ms | 否 | ✓ 1967ms | http |
| 116.80.96.95:3172 | ✓ 1598ms | 否 | 否 | ✓ 1855ms | ✓ 1662ms | http |
| 138.124.90.140:1080 | ✓ 1099ms | 否 | ✓ 1365ms | ✓ 1752ms | ✓ 1347ms | http |
| 103.183.10.172:3125 | 否 | 否 | ✓ 1286ms | ✓ 1428ms | ✓ 1392ms | http |
| 111.228.59.13:8888 | ✓ 1446ms | ✓ 1389ms | ✓ 1436ms | ✓ 1990ms | 否 | http |
| 36.37.180.40:8080 | ✓ 1421ms | 否 | ✓ 1842ms | ✓ 1628ms | ✓ 1552ms | http |

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
