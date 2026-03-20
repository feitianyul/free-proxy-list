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

最后更新：2026-03-20 12:25:57 UTC（2026-03-20 20:25:57 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 178.156.187.185:10001 | ✓ 444ms | ✓ 995ms | ✓ 217ms | ✓ 1121ms | ✓ 993ms | http |
| 147.161.210.140:8800 | ✓ 1620ms | 否 | ✓ 1121ms | ✓ 1877ms | ✓ 1162ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1572ms | ✓ 1469ms | ✓ 787ms | http |
| 113.160.132.26:8080 | ✓ 1786ms | 否 | ✓ 1056ms | 否 | ✓ 1136ms | http |
| 103.166.185.54:3128 | ✓ 1790ms | ✓ 1824ms | ✓ 1342ms | ✓ 1851ms | 否 | http |
| 45.167.124.52:8080 | 否 | ✓ 1805ms | ✓ 1428ms | 否 | ✓ 1592ms | http |
| 1.231.81.166:3128 | ✓ 1667ms | ✓ 1596ms | 否 | ✓ 1719ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 903ms | ✓ 933ms | ✓ 1020ms | http |
| 137.220.150.22:6005 | ✓ 873ms | 否 | ✓ 893ms | ✓ 1290ms | ✓ 1003ms | http |
| 167.103.34.108:8800 | ✓ 1265ms | 否 | ✓ 1308ms | ✓ 1463ms | ✓ 1390ms | http |
| 167.103.31.122:8800 | ✓ 1685ms | 否 | ✓ 1703ms | ✓ 1549ms | ✓ 1891ms | http |
| 46.101.190.71:3128 | ✓ 522ms | 否 | ✓ 1729ms | ✓ 1755ms | ✓ 1267ms | http |
| 162.240.154.26:3128 | ✓ 425ms | ✓ 1875ms | 否 | ✓ 1932ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1079ms | 否 | ✓ 1107ms | ✓ 1592ms | 否 | http |
| 137.220.150.170:6005 | ✓ 1332ms | 否 | 否 | ✓ 1515ms | ✓ 1116ms | http |
| 194.147.115.50:3128 | ✓ 890ms | ✓ 1570ms | ✓ 689ms | ✓ 1877ms | ✓ 1278ms | http |
| 147.161.239.240:8800 | ✓ 885ms | 否 | ✓ 1064ms | ✓ 1527ms | ✓ 1492ms | http |
| 154.64.243.50:7890 | ✓ 1409ms | 否 | ✓ 563ms | 否 | ✓ 1752ms | http |
| 192.71.213.85:9812 | ✓ 1447ms | 否 | ✓ 1683ms | ✓ 1618ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1115ms | ✓ 1429ms | 否 | ✓ 1647ms | ✓ 1136ms | http |
| 149.62.191.202:3128 | ✓ 1451ms | ✓ 1968ms | ✓ 1717ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 1866ms | ✓ 1284ms | 否 | ✓ 1353ms | ✓ 1073ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1417ms | ✓ 985ms | ✓ 1658ms | ✓ 1361ms | http |
| 46.246.1.106:3128 | 否 | ✓ 1611ms | ✓ 910ms | ✓ 1555ms | ✓ 1116ms | http |
| 137.220.150.152:6005 | ✓ 1845ms | 否 | ✓ 1118ms | ✓ 1465ms | ✓ 1832ms | http |
| 91.238.105.64:2024 | ✓ 1048ms | ✓ 1950ms | ✓ 1844ms | ✓ 1542ms | ✓ 1200ms | http |
| 88.80.150.82:8080 | ✓ 1418ms | ✓ 1907ms | 否 | 否 | ✓ 1259ms | https |
| 147.45.67.148:8080 | ✓ 1118ms | 否 | ✓ 1383ms | 否 | ✓ 1279ms | http |
| 84.247.149.172:3128 | ✓ 853ms | 否 | ✓ 1696ms | ✓ 1215ms | ✓ 942ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1308ms | ✓ 1610ms | ✓ 1467ms | http |
| 106.75.15.167:7890 | 否 | ✓ 1291ms | ✓ 1003ms | ✓ 1300ms | 否 | http |
| 137.220.151.110:6005 | ✓ 1415ms | 否 | ✓ 1671ms | ✓ 1708ms | ✓ 1247ms | http |
| 172.212.68.37:3128 | ✓ 212ms | 否 | ✓ 790ms | ✓ 1380ms | ✓ 1118ms | http |
| 111.201.98.211:7890 | ✓ 1766ms | ✓ 1282ms | 否 | 否 | ✓ 1768ms | http |
| 106.117.208.101:7890 | ✓ 1169ms | 否 | ✓ 1148ms | 否 | ✓ 1091ms | http |
| 43.153.9.19:80 | ✓ 791ms | ✓ 1399ms | ✓ 882ms | ✓ 1353ms | ✓ 1294ms | http |
| 177.247.249.5:3128 | ✓ 1288ms | 否 | ✓ 1899ms | 否 | ✓ 1392ms | http |
| 133.242.138.34:8100 | 否 | ✓ 1201ms | ✓ 735ms | ✓ 1617ms | 否 | http |
| 137.220.150.104:6005 | ✓ 1307ms | 否 | ✓ 1015ms | ✓ 1411ms | ✓ 967ms | http |
| 103.154.214.50:3128 | ✓ 1704ms | 否 | ✓ 945ms | ✓ 1312ms | ✓ 1018ms | http |
| 45.129.141.143:3128 | ✓ 1141ms | 否 | ✓ 1511ms | ✓ 1959ms | ✓ 1488ms | http |
| 38.180.2.107:3128 | ✓ 1420ms | ✓ 1971ms | ✓ 1980ms | 否 | ✓ 1921ms | http |
| 185.191.236.162:3128 | ✓ 1104ms | 否 | ✓ 1913ms | 否 | ✓ 1808ms | http |
| 142.171.224.229:7890 | ✓ 382ms | ✓ 881ms | ✓ 800ms | ✓ 818ms | 否 | http |
| 102.134.49.165:6005 | ✓ 788ms | ✓ 1509ms | 否 | 否 | ✓ 1194ms | http |
| 101.47.73.135:3128 | ✓ 1366ms | 否 | ✓ 1719ms | ✓ 1384ms | ✓ 1770ms | http |
| 103.39.51.190:8080 | ✓ 1942ms | 否 | 否 | ✓ 1554ms | ✓ 1511ms | http |
| 8.212.172.106:8080 | ✓ 1099ms | 否 | ✓ 1213ms | ✓ 1505ms | ✓ 1122ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1198ms | ✓ 1460ms | ✓ 1839ms | http |
| 137.184.6.37:3128 | ✓ 1940ms | ✓ 1381ms | ✓ 403ms | ✓ 899ms | ✓ 683ms | http |
| 45.140.147.155:1081 | 否 | ✓ 1115ms | ✓ 1880ms | ✓ 1898ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1960ms | 否 | ✓ 1689ms | 否 | ✓ 1729ms | http |
| 38.145.218.232:8448 | 否 | 否 | ✓ 255ms | ✓ 859ms | ✓ 1217ms | http |
| 38.145.218.218:8449 | 否 | ✓ 1687ms | ✓ 476ms | ✓ 882ms | ✓ 1033ms | http |
| 38.34.179.191:8452 | 否 | ✓ 1795ms | ✓ 1224ms | ✓ 902ms | ✓ 1010ms | http |
| 121.230.9.184:1080 | ✓ 1870ms | ✓ 1537ms | ✓ 1431ms | 否 | ✓ 1237ms | http |
| 180.103.19.252:1080 | ✓ 1269ms | 否 | ✓ 1322ms | 否 | ✓ 1370ms | http |
| 45.136.131.38:8449 | 否 | 否 | ✓ 1950ms | ✓ 1115ms | ✓ 875ms | http |
| 45.149.92.147:5001 | ✓ 745ms | 否 | ✓ 749ms | ✓ 1161ms | 否 | http |
| 38.145.208.192:8448 | ✓ 862ms | ✓ 1352ms | ✓ 1432ms | ✓ 886ms | ✓ 916ms | http |
| 45.136.131.44:8452 | ✓ 845ms | ✓ 1630ms | 否 | ✓ 1148ms | ✓ 1658ms | http |
| 133.18.110.87:1081 | 否 | ✓ 1545ms | ✓ 842ms | ✓ 1095ms | 否 | http |
| 38.55.104.182:6005 | ✓ 1841ms | 否 | ✓ 1272ms | ✓ 1071ms | ✓ 858ms | http |
| 150.249.255.91:3128 | ✓ 1513ms | 否 | ✓ 1558ms | ✓ 1128ms | 否 | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1972ms | ✓ 1865ms | ✓ 1987ms | http |

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
