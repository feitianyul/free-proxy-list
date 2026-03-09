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

最后更新：2026-03-09 09:42:23 UTC（2026-03-09 17:42:23 UTC+8）

**代理总数：63**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 64.186.232.4:10808 | ✓ 629ms | ✓ 712ms | ✓ 415ms | ✓ 838ms | ✓ 759ms | http |
| 1.231.81.166:3128 | ✓ 1942ms | 否 | ✓ 1064ms | ✓ 1079ms | ✓ 878ms | http |
| 101.47.73.135:3128 | ✓ 1737ms | 否 | 否 | ✓ 1723ms | ✓ 1654ms | http |
| 47.77.193.180:1080 | ✓ 232ms | ✓ 897ms | ✓ 404ms | ✓ 849ms | ✓ 624ms | http |
| 115.231.181.40:8128 | ✓ 1082ms | ✓ 1350ms | ✓ 1111ms | ✓ 1585ms | ✓ 1155ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1291ms | ✓ 1902ms | ✓ 1571ms | http |
| 116.80.49.168:3172 | ✓ 1418ms | 否 | ✓ 1667ms | 否 | ✓ 1739ms | http |
| 35.225.22.61:80 | ✓ 748ms | ✓ 1187ms | 否 | ✓ 1053ms | ✓ 818ms | http |
| 194.213.18.200:443 | ✓ 1211ms | ✓ 1690ms | 否 | 否 | ✓ 1265ms | http |
| 154.64.240.39:1080 | ✓ 1677ms | 否 | 否 | ✓ 1635ms | ✓ 1074ms | http |
| 162.248.165.72:1080 | ✓ 1102ms | 否 | 否 | ✓ 1436ms | ✓ 1977ms | http |
| 101.43.255.96:80 | ✓ 1092ms | ✓ 1409ms | ✓ 1173ms | 否 | ✓ 1157ms | http |
| 81.70.169.194:80 | ✓ 1536ms | 否 | ✓ 1038ms | ✓ 1425ms | ✓ 1159ms | http |
| 165.227.5.10:8888 | ✓ 647ms | 否 | ✓ 932ms | ✓ 1965ms | 否 | http |
| 67.169.98.211:443 | ✓ 1472ms | 否 | ✓ 1787ms | ✓ 1866ms | 否 | http |
| 136.49.34.18:8888 | ✓ 1220ms | 否 | ✓ 1121ms | 否 | ✓ 1907ms | http |
| 116.80.49.166:3172 | ✓ 1709ms | 否 | ✓ 1800ms | ✓ 1959ms | ✓ 1816ms | http |
| 46.39.105.157:8080 | ✓ 630ms | ✓ 1601ms | ✓ 1713ms | 否 | 否 | http |
| 159.223.42.219:3128 | ✓ 943ms | 否 | ✓ 1999ms | ✓ 1163ms | ✓ 934ms | http |
| 168.235.110.63:3128 | ✓ 1253ms | 否 | 否 | ✓ 1065ms | ✓ 785ms | http |
| 41.33.219.130:1981 | ✓ 1923ms | 否 | 否 | ✓ 1859ms | ✓ 1588ms | http |
| 121.237.181.137:8888 | ✓ 1010ms | ✓ 1237ms | ✓ 1003ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1093ms | ✓ 1351ms | ✓ 1329ms | 否 | ✓ 1102ms | http |
| 120.92.212.16:7890 | ✓ 1936ms | 否 | 否 | ✓ 1342ms | ✓ 1899ms | http |
| 107.172.125.217:3128 | ✓ 864ms | 否 | ✓ 958ms | ✓ 901ms | ✓ 709ms | http |
| 62.113.119.14:8080 | ✓ 1276ms | ✓ 1563ms | ✓ 1006ms | ✓ 1560ms | ✓ 1136ms | http |
| 178.236.245.17:3128 | ✓ 1269ms | 否 | ✓ 988ms | ✓ 1817ms | ✓ 1567ms | http |
| 103.139.138.194:3128 | ✓ 1779ms | 否 | ✓ 1208ms | ✓ 1526ms | ✓ 1490ms | http |
| 178.236.245.59:3128 | ✓ 1269ms | ✓ 1849ms | ✓ 969ms | ✓ 1919ms | 否 | http |
| 190.9.109.207:999 | ✓ 861ms | 否 | ✓ 875ms | ✓ 1226ms | ✓ 988ms | http |
| 190.9.109.198:999 | ✓ 757ms | ✓ 1489ms | ✓ 1248ms | ✓ 1233ms | 否 | http |
| 39.104.201.40:7890 | ✓ 1437ms | ✓ 1348ms | 否 | ✓ 1372ms | 否 | http |
| 116.80.49.170:3172 | ✓ 1873ms | 否 | ✓ 1706ms | ✓ 1990ms | 否 | http |
| 116.80.49.169:3172 | ✓ 1875ms | 否 | ✓ 1702ms | 否 | ✓ 1841ms | http |
| 116.80.49.167:3172 | ✓ 1878ms | 否 | ✓ 1692ms | ✓ 1976ms | 否 | http |
| 116.80.49.162:3172 | ✓ 1726ms | 否 | ✓ 1669ms | ✓ 1992ms | ✓ 1778ms | http |
| 116.80.49.156:3172 | ✓ 1731ms | 否 | ✓ 1667ms | 否 | ✓ 1770ms | http |
| 116.80.49.159:3172 | ✓ 1732ms | 否 | ✓ 1666ms | ✓ 1974ms | 否 | http |
| 116.80.49.172:3172 | ✓ 1695ms | 否 | ✓ 1667ms | 否 | ✓ 1768ms | http |
| 116.80.49.165:3172 | ✓ 1723ms | 否 | ✓ 1669ms | 否 | ✓ 1780ms | http |
| 45.136.198.40:3128 | ✓ 1248ms | ✓ 1947ms | 否 | ✓ 1989ms | ✓ 1629ms | http |
| 103.166.185.54:3128 | ✓ 1445ms | ✓ 1729ms | ✓ 1027ms | ✓ 1337ms | ✓ 1087ms | http |
| 118.113.246.172:1080 | ✓ 1396ms | 否 | ✓ 1423ms | 否 | ✓ 1549ms | http |
| 103.155.167.82:8082 | ✓ 1853ms | 否 | 否 | ✓ 1537ms | ✓ 1478ms | http |
| 1.225.116.115:1080 | 否 | 否 | ✓ 1009ms | ✓ 1210ms | ✓ 1236ms | http |
| 116.80.49.163:3172 | ✓ 1904ms | 否 | 否 | ✓ 1953ms | ✓ 1802ms | http |
| 116.80.49.161:3172 | ✓ 1903ms | 否 | ✓ 1636ms | ✓ 1969ms | ✓ 1787ms | http |
| 103.82.23.118:5185 | ✓ 1855ms | 否 | ✓ 1312ms | 否 | ✓ 1596ms | http |
| 207.254.71.62:8088 | ✓ 986ms | 否 | ✓ 1480ms | ✓ 1857ms | ✓ 1546ms | http |
| 45.129.141.143:3128 | ✓ 1729ms | 否 | ✓ 1433ms | 否 | ✓ 1701ms | http |
| 172.212.68.37:3128 | ✓ 1314ms | ✓ 1518ms | ✓ 1310ms | ✓ 1902ms | ✓ 1124ms | http |
| 45.186.6.104:3128 | ✓ 1877ms | ✓ 1807ms | ✓ 1657ms | 否 | 否 | http |
| 164.90.151.28:3128 | ✓ 428ms | ✓ 1560ms | ✓ 1347ms | ✓ 881ms | ✓ 668ms | http |
| 47.95.231.180:8084 | 否 | ✓ 1321ms | ✓ 1055ms | ✓ 1318ms | ✓ 1050ms | http |
| 34.101.184.164:3128 | ✓ 1646ms | 否 | ✓ 1825ms | ✓ 1935ms | ✓ 1106ms | http |
| 45.140.147.82:1081 | ✓ 1007ms | ✓ 1129ms | ✓ 1436ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1816ms | 否 | 否 | ✓ 1647ms | ✓ 1556ms | http |
| 192.71.213.85:5678 | ✓ 965ms | 否 | ✓ 1853ms | ✓ 1934ms | 否 | http |
| 192.71.213.85:9091 | ✓ 890ms | 否 | ✓ 519ms | ✓ 1553ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1027ms | ✓ 1943ms | 否 | 否 | ✓ 1802ms | https |
| 194.67.105.229:3128 | ✓ 1172ms | 否 | ✓ 1815ms | 否 | ✓ 1731ms | http |
| 113.132.112.110:9000 | ✓ 1551ms | ✓ 1691ms | 否 | 否 | ✓ 1860ms | http |
| 205.209.118.30:3138 | ✓ 396ms | ✓ 996ms | 否 | ✓ 1161ms | ✓ 1872ms | http |

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
